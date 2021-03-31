package client

import (
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
	grpcConn    *grpc.ClientConn
	storeClient pb.StorerClient
}

func (s *Store) Close() error {
	return s.grpcConn.Close()
}

func (s *Store) NewEncodedObject() plumbing.EncodedObject {
	panic("implement me")
}

func (s *Store) SetEncodedObject(object plumbing.EncodedObject) (plumbing.Hash, error) {
	panic("implement me")
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
