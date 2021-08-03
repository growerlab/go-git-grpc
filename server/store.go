package server

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/common"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

var (
	ErrNotFoundObject = errors.New("Not found object")
	ErrNotFoundIter   = errors.New("Not found iter")
)

var _ pb.StorerServer = (*Store)(nil)

func NewStore(root string) *Store {
	return &Store{
		root:        root,
		objectStash: NewObjectCache(0),
		iterStash:   NewObjectCache(0),
	}
}

type Store struct {
	*pb.UnimplementedStorerServer

	// 仓库根目录
	root string

	// 对象缓存
	objectStash *ObjectCache
	// 迭代器缓存
	iterStash *ObjectCache
}

func (s *Store) SetShallow(ctx context.Context, hashs *pb.Hashs) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, hashs.RepoPath, func(r *git.Repository) error {
		commits := make([]plumbing.Hash, len(hashs.Hash))
		for i := range hashs.Hash {
			commits[i] = plumbing.NewHash(hashs.Hash[i])
		}
		err := r.Storer.SetShallow(commits)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) Shallow(ctx context.Context, none *pb.None) (*pb.Hashs, error) {
	var result = new(pb.Hashs)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		commits, err := r.Storer.Shallow()
		if err != nil {
			return errors.WithStack(err)
		}
		for _, cmt := range commits {
			result.Hash = append(result.Hash, cmt.String())
		}
		return nil
	})
	return result, err
}

func (s *Store) SetIndex(ctx context.Context, idx *pb.Index) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, idx.RepoPath, func(r *git.Repository) error {
		newIdx := common.BuildPbRefToIndex(idx)
		err := r.Storer.SetIndex(newIdx)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) GetIndex(ctx context.Context, none *pb.None) (*pb.Index, error) {
	var result *pb.Index
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		idx, err := r.Storer.Index()
		if err != nil {
			return errors.WithStack(err)
		}
		result = common.BuildIndexToPbRef(idx)
		return nil
	})
	return result, err
}

func (s *Store) GetConfig(ctx context.Context, none *pb.None) (*pb.Config, error) {
	var result = new(pb.Config)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		cfg, err := r.Storer.Config()
		if err != nil {
			return errors.WithStack(err)
		}
		result = common.BuildConfigFromPbConfig(cfg)
		return nil
	})
	return result, err
}

func (s *Store) SetConfig(ctx context.Context, c *pb.Config) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, c.RepoPath, func(r *git.Repository) error {
		cfg := common.BuildPbConfigFromConfig(c)
		err := r.Storer.SetConfig(cfg)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) Modules(ctx context.Context, none *pb.None) (*pb.ModuleNames, error) {
	var result = new(pb.ModuleNames)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		cfg, err := r.Storer.Config()
		if err != nil {
			return errors.WithStack(err)
		}
		for _, submd := range cfg.Submodules {
			result.Names = append(result.Names, submd.Name)
		}
		return nil
	})
	return result, err
}

func (s *Store) mustEmbedUnimplementedStorerServer() {}
