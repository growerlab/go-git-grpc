package client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	format "github.com/go-git/go-git/v5/plumbing/format/config"
	"github.com/go-git/go-git/v5/plumbing/format/index"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage"
	"github.com/growerlab/go-git-grpc/common"
	"github.com/growerlab/go-git-grpc/pb"
	"google.golang.org/grpc"
)

var _ storage.Storer = (*Store)(nil)

func NewStore(ctx context.Context, repPath string, grpcConn *grpc.ClientConn, pbClient pb.StorerClient) *Store {
	return &Store{
		repoPath: repPath,
		lastErr:  nil,
		ctx:      ctx,
		grpcConn: grpcConn,
		client:   pbClient,
	}
}

type Store struct {
	repoPath string
	lastErr  error

	ctx      context.Context
	grpcConn *grpc.ClientConn
	client   pb.StorerClient
}

func (s *Store) Close() error {
	err := s.grpcConn.Close()
	return err
}

func (s *Store) NewEncodedObject() plumbing.EncodedObject {
	params := &pb.None{RepoPath: s.repoPath}
	resp, err := s.client.NewEncodedObject(s.ctx, params)
	if err != nil {
		s.lastErr = errors.WithStack(err)
		return nil
	}

	return &EncodedObject{
		ctx:      s.ctx,
		client:   s.client,
		repoPath: s.repoPath,
		uuid:     resp.Value,
	}
}

func (s *Store) SetEncodedObject(obj plumbing.EncodedObject) (plumbing.Hash, error) {
	ob := obj.(*EncodedObject)
	params := &pb.UUID{Value: ob.uuid}

	hash, err := s.client.SetEncodedObject(s.ctx, params)
	if err != nil {
		return plumbing.ZeroHash, errors.WithStack(err)
	}
	return plumbing.NewHash(hash.Value), nil
}

func (s *Store) EncodedObject(objectType plumbing.ObjectType, hash plumbing.Hash) (plumbing.EncodedObject, error) {
	params := &pb.GetEncodeObject{
		RepoPath: s.repoPath,
		Hash:     hash.String(),
		Type:     objectType.String(),
	}
	obj, err := s.client.EncodedObjectEntity(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := &EncodedObject{
		ctx:            s.ctx,
		client:         s.client,
		repoPath:       s.repoPath,
		uuid:           obj.UUID,
		readonlyObject: buildReadonlyEncodedObject(obj),
	}
	return result, nil
}

func (s *Store) IterEncodedObjects(objectType plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	panic("implement me")
}

func (s *Store) HasEncodedObject(hash plumbing.Hash) error {
	panic("implement me")
}

func (s *Store) EncodedObjectSize(hash plumbing.Hash) (int64, error) {
	panic("implement me")
}

func (s *Store) SetReference(reference *plumbing.Reference) error {
	panic("implement me")
}

func (s *Store) CheckAndSetReference(new, old *plumbing.Reference) error {
	panic("implement me")
}

func (s *Store) Reference(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	params := &pb.ReferenceName{
		RepoPath: s.repoPath,
		Name:     name.String(),
	}
	result, err := s.client.GetReference(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return plumbing.NewReferenceFromStrings(result.N, result.Target), nil
}

func (s *Store) IterReferences() (storer.ReferenceIter, error) {
	params := &pb.None{
		RepoPath: s.repoPath,
	}
	pbRefs, err := s.client.GetReferences(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	refs := make([]*plumbing.Reference, 0, len(pbRefs.Refs))

	for _, r := range pbRefs.Refs {
		ref := plumbing.NewReferenceFromStrings(r.N, r.Target)
		refs = append(refs, ref)
	}

	return storer.NewReferenceSliceIter(refs), nil
}

func (s *Store) RemoveReference(name plumbing.ReferenceName) error {
	params := &pb.ReferenceName{
		RepoPath: s.repoPath,
		Name:     string(name),
	}
	_, err := s.client.RemoveReference(s.ctx, params)
	return errors.WithStack(err)
}

func (s *Store) CountLooseRefs() (int, error) {
	params := &pb.None{
		RepoPath: s.repoPath,
	}
	n, err := s.client.CountLooseRefs(s.ctx, params)
	return int(n.Value), errors.WithStack(err)
}

func (s *Store) PackRefs() error {
	params := &pb.None{
		RepoPath: s.repoPath,
	}
	_, err := s.client.PackRefs(s.ctx, params)
	return errors.WithStack(err)
}

func (s *Store) SetShallow(hashes []plumbing.Hash) error {
	if len(hashes) == 0 {
		return nil
	}
	hashesStrs := make([]string, 0, len(hashes))
	for _, h := range hashes {
		hashesStrs = append(hashesStrs, h.String())
	}

	params := &pb.Hashs{
		RepoPath: s.repoPath,
		Hash:     hashesStrs,
	}
	_, err := s.client.SetShallow(s.ctx, params)
	return errors.WithStack(err)
}

func (s *Store) Shallow() ([]plumbing.Hash, error) {
	params := &pb.None{
		RepoPath: s.repoPath,
	}
	rawHashes, err := s.client.Shallow(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]plumbing.Hash, 0, len(rawHashes.Hash))
	for _, h := range rawHashes.Hash {
		result = append(result, plumbing.NewHash(h))
	}
	return result, nil
}

func (s *Store) SetIndex(index *index.Index) error {
	params := &pb.Index{
		RepoPath: s.repoPath,
	}
	_, err := s.client.SetIndex(s.ctx, params)
	return errors.WithStack(err)
}

func (s *Store) Index() (*index.Index, error) {
	params := &pb.None{
		RepoPath: s.repoPath,
	}
	idx, err := s.client.GetIndex(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return common.BuildPbRefToIndex(idx), nil
}

func (s *Store) Config() (*config.Config, error) {
	params := &pb.None{
		RepoPath: s.repoPath,
		UUID:     "",
	}
	cfg, err := s.client.GetConfig(s.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	remotes := map[string]*config.RemoteConfig{}
	submodules := map[string]*config.Submodule{}
	branches := map[string]*config.Branch{}
	raw := new(format.Config)
	err = json.Unmarshal(cfg.Raw, raw)
	if err != nil {
		log.Printf("Config() unmarshal was err:%+v\n", err)
	}

	for _, r := range cfg.Remotes {
		var fetches = make([]config.RefSpec, 0, len(r.Config.Fetch))
		for _, f := range r.Config.Fetch {
			fetches = append(fetches, config.RefSpec(f))
		}
		remotes[r.Key] = &config.RemoteConfig{
			Name:  r.Config.Name,
			URLs:  r.Config.URLs,
			Fetch: fetches,
		}
	}

	for _, sub := range cfg.Submodules {
		s := &config.Submodule{
			Name:   sub.Sub.Name,
			Path:   sub.Sub.Path,
			URL:    sub.Sub.URL,
			Branch: sub.Sub.Branch,
		}
		submodules[sub.Key] = s
	}

	for _, bn := range cfg.Branches {
		b := &config.Branch{
			Name:   bn.Branch.Name,
			Remote: bn.Branch.Remote,
			Merge:  plumbing.NewBranchReferenceName(bn.Branch.Merge),
			Rebase: bn.Branch.Rebase,
		}
		branches[bn.Key] = b
	}

	return &config.Config{
		Core: struct {
			IsBare      bool
			Worktree    string
			CommentChar string
		}{
			IsBare:      cfg.Core.IsBare,
			Worktree:    cfg.Core.Worktree,
			CommentChar: cfg.Core.CommentChar,
		},
		User: struct {
			Name  string
			Email string
		}{
			Name:  cfg.User.Name,
			Email: cfg.User.Email,
		},
		Author: struct {
			Name  string
			Email string
		}{
			Name:  cfg.Author.Name,
			Email: cfg.Author.Email,
		},
		Committer: struct {
			Name  string
			Email string
		}{
			Name:  cfg.Committer.Name,
			Email: cfg.Committer.Email,
		},
		Pack: struct {
			Window uint
		}{
			Window: uint(cfg.Pack.Window),
		},
		Remotes:    remotes,
		Submodules: submodules,
		Branches:   branches,
		Raw:        raw,
	}, nil
}

func (s *Store) SetConfig(config *config.Config) error {
	panic("implement me")
}

func (s *Store) Module(name string) (storage.Storer, error) {
	panic("implement me")
}

func buildReadonlyEncodedObject(obj *pb.EncodedObject) plumbing.EncodedObject {
	typ, _ := plumbing.ParseObjectType(obj.Type)
	return &ReadonlyEncodedObject{
		hash: plumbing.NewHash(obj.Hash),
		typ:  typ,
		size: obj.GetSize(),
	}
}
