package server

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

func (s *Store) NewEncodedObjectIter(ctx context.Context, tp *pb.ObjectType) (*pb.UUID, error) {
	var (
		uuid     = buildUUID()
		result   = &pb.UUID{Value: uuid}
		repoPath = tp.RepoPath
	)

	err := repo(s.root, repoPath, func(repo *git.Repository) error {
		objType, err := plumbing.ParseObjectType(tp.Type)
		if err != nil {
			return errors.WithStack(err)
		}
		iter, err := repo.Storer.IterEncodedObjects(objType)
		if err != nil {
			return errors.WithStack(err)
		}

		s.putIter(&EncodedObjectIterExt{
			EncodedObjectIter: iter,
			uuid:              uuid,
		})
		return nil
	})
	return result, err
}

func (s *Store) EncodedObjectNext(ctx context.Context, none *pb.None) (*pb.EncodedObject, error) {
	var (
		iterUUID = none.UUID
		repoPath = none.RepoPath
	)

	iter, ok := s.getIter(iterUUID)
	if !ok {
		return nil, ErrNotFoundIter
	}
	obj, err := iter.Next()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	newObj := NewEncodedObject(ctx, buildUUID(), repoPath, obj)
	s.putObject(newObj)

	return newObj.PBEncodeObject(), nil
}

func (s *Store) EncodedObjectForEach(none *pb.None, stream pb.Storer_EncodedObjectForEachServer) error {
	var (
		iterUUID = none.UUID
		repoPath = none.RepoPath
	)
	iter, ok := s.getIter(iterUUID)
	if !ok {
		return ErrNotFoundIter
	}

	err := iter.ForEach(func(object plumbing.EncodedObject) error {
		obj := NewEncodedObject(context.Background(), buildUUID(), repoPath, object)
		s.putObject(obj)
		return stream.Send(obj.PBEncodeObject())
	})
}

func (s *Store) putIter(iter *EncodedObjectIterExt) {
	s.iterStash.Set(iter)
}
func (s *Store) getIter(uuid string) (*EncodedObjectIterExt, bool) {
	iterObj, ok := s.iterStash.Get(uuid)
	if !ok {
		return nil, false
	}
	iter, ok := iterObj.(*EncodedObjectIterExt)
	return iter, ok
}
