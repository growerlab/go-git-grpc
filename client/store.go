package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/index"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage"
	"github.com/growerlab/go-git-grpc/pb"
	"google.golang.org/grpc"
)

var _ storage.Storer = (*Store)(nil)

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
	panic("implement me")
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
	panic("implement me")
}

func (s *Store) IterReferences() (storer.ReferenceIter, error) {
	panic("implement me")
}

func (s *Store) RemoveReference(name plumbing.ReferenceName) error {
	panic("implement me")
}

func (s *Store) CountLooseRefs() (int, error) {
	panic("implement me")
}

func (s *Store) PackRefs() error {
	panic("implement me")
}

func (s *Store) SetShallow(hashes []plumbing.Hash) error {
	panic("implement me")
}

func (s *Store) Shallow() ([]plumbing.Hash, error) {
	panic("implement me")
}

func (s *Store) SetIndex(index *index.Index) error {
	panic("implement me")
}

func (s *Store) Index() (*index.Index, error) {
	panic("implement me")
}

func (s *Store) Config() (*config.Config, error) {
	panic("implement me")
}

func (s *Store) SetConfig(config *config.Config) error {
	panic("implement me")
}

func (s *Store) Module(name string) (storage.Storer, error) {
	panic("implement me")
}
