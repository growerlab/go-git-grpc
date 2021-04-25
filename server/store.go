package server

import (
	"bytes"
	"context"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/common"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

var (
	ErrNotFoundObject = errors.New("Not found object")
)

var _ pb.StorerServer = (*Store)(nil)

func NewStore(root string) *Store {
	return &Store{
		root: root,
	}
}

type Store struct {
	*pb.UnimplementedStorerServer

	// 仓库根目录
	root string

	objectLRU *ObjectLRU
}

func (s *Store) NewEncodedObject(ctx context.Context, none *pb.None) (*pb.UUID, error) {
	obj := NewEncodedObject(ctx, buildUUID(), none.RepoPath, &plumbing.MemoryObject{})
	s.objectLRU.Put(obj)
	return &pb.UUID{Value: obj.UUID()}, nil
}

func (s *Store) SetEncodedObject(ctx context.Context, uuid *pb.UUID) (*pb.Hash, error) {
	var result *pb.Hash
	var key = uuid.Value
	var obj, exists = s.getObject(key)
	if !exists {
		return nil, ErrNotFoundObject
	}

	err := repo(s.root, obj.repoPath, func(r *git.Repository) error {
		h, err := r.Storer.SetEncodedObject(obj)
		if err != nil {
			return errors.WithStack(err)
		}
		result = &pb.Hash{
			Value: h.String(),
		}
		return nil
	})
	return result, errors.WithStack(err)
}
func (s *Store) SetEncodedObjectType(ctx context.Context, i *pb.Int) (*pb.None, error) {
	var result = &pb.None{UUID: i.UUID}
	var objectType = plumbing.ObjectType(i.Value)

	obj, exists := s.getObject(i.UUID)
	if !exists {
		return nil, ErrNotFoundObject
	}
	obj.SetType(objectType)

	return result, nil
}

func (s *Store) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
	var result = &pb.None{UUID: i.UUID}

	obj, exists := s.getObject(i.UUID)
	if !exists {
		return nil, ErrNotFoundObject
	}
	obj.SetSize(i.Value)

	return result, nil
}

func (s *Store) EncodedObjectEntity(ctx context.Context, objEntity *pb.GetEncodeObject) (*pb.EncodedObject, error) {
	var result *pb.EncodedObject
	var repoPath = objEntity.RepoPath

	err := repo(s.root, objEntity.RepoPath, func(r *git.Repository) error {
		objectType, err := plumbing.ParseObjectType(objEntity.Type)
		if err != nil {
			return errors.WithStack(err)
		}
		hash := plumbing.NewHash(objEntity.Hash)
		obj, err := r.Storer.EncodedObject(objectType, hash)
		if err != nil {
			return err
		}

		newObj := NewEncodedObject(ctx, buildUUID(), repoPath, obj)
		s.objectLRU.Put(newObj)

		result = newObj.PBEncodeObject()
		return nil
	})
	return result, err
}

func (s *Store) EncodedObjectType(ctx context.Context, none *pb.None) (*pb.Int, error) {
	obj, exists := s.getObject(none.UUID)
	if !exists {
		return nil, ErrNotFoundObject
	}
	return &pb.Int{Value: int32(obj.Type())}, nil
}

func (s *Store) EncodedObjectHash(ctx context.Context, none *pb.None) (*pb.Hash, error) {
	obj, exists := s.getObject(none.UUID)
	if !exists {
		return nil, ErrNotFoundObject
	}
	return &pb.Hash{Value: obj.uuid}, nil
}

func (s *Store) EncodedObjectSize(ctx context.Context, none *pb.None) (*pb.Int64, error) {
	obj, exists := s.getObject(none.UUID)
	if !exists {
		return nil, ErrNotFoundObject
	}
	return &pb.Int64{Value: obj.Size()}, nil
}

func (s *Store) EncodedObjectRWStream(stream pb.Storer_EncodedObjectRWStreamServer) error {
	first, err := stream.Recv()
	if err != nil {
		return errors.WithStack(err)
	}
	key := first.UUID
	obj, exists := s.getObject(key)
	if !exists {
		return ErrNotFoundObject
	}

	switch first.Flag {
	case pb.RWStream_READ:
		reader, err := obj.Reader()
		if err != nil {
			return errors.WithStack(err)
		}
		defer reader.Close()

		var buf = make([]byte, bytes.MinRead)
		for {
			var n int
			n, err = reader.Read(buf)
			if err == io.EOF {
				return err
			}
			buf = buf[:n]
			err = stream.Send(&pb.RWStream{
				Value: buf,
			})
			if err != nil {
				return err
			}
			buf = buf[:bytes.MinRead]
		}
	case pb.RWStream_WRITE:
		writer, err := obj.Writer()
		if err != nil {
			return errors.WithStack(err)
		}
		defer writer.Close()

		for {
			rw, err := stream.Recv()
			if err == io.EOF {
				return err
			}
			_, err = writer.Write(rw.Value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) SetReference(ctx context.Context, reference *pb.Reference) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, reference.RepoPath, func(r *git.Repository) error {
		ref := plumbing.NewReferenceFromStrings(reference.N, reference.Target)
		err := r.Storer.SetReference(ref)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) CheckAndSetReference(ctx context.Context, params *pb.SetReferenceParams) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, params.RepoPath, func(r *git.Repository) error {
		newRef := plumbing.NewReferenceFromStrings(params.New.N, params.New.Target)
		oldRef := plumbing.NewReferenceFromStrings(params.Old.N, params.Old.Target)
		err := r.Storer.CheckAndSetReference(newRef, oldRef)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) GetReference(ctx context.Context, name *pb.ReferenceName) (*pb.Reference, error) {
	var result *pb.Reference
	err := repo(s.root, name.RepoPath, func(r *git.Repository) error {
		ref, err := r.Storer.Reference(plumbing.ReferenceName(name.Name))
		if err != nil {
			return errors.WithStack(err)
		}

		result = common.BuildRefToPbRef(ref)
		return nil
	})
	return result, err
}

func (s *Store) GetReferences(ctx context.Context, none *pb.None) (*pb.References, error) {
	var result = new(pb.References)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		iter, err := r.Storer.IterReferences()
		if err != nil {
			return errors.WithStack(err)
		}

		err = iter.ForEach(func(ref *plumbing.Reference) error {
			pbRef := common.BuildRefToPbRef(ref)
			result.Refs = append(result.Refs, pbRef)
			return nil
		})
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) RemoveReference(ctx context.Context, name *pb.ReferenceName) (*pb.None, error) {
	err := repo(s.root, name.RepoPath, func(r *git.Repository) error {
		rn := plumbing.ReferenceName(name.Name)
		err := r.Storer.RemoveReference(rn)
		return errors.WithStack(err)
	})
	return &pb.None{}, err
}

func (s *Store) CountLooseRefs(ctx context.Context, none *pb.None) (*pb.Int64, error) {
	var result = new(pb.Int64)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		count, err := r.Storer.CountLooseRefs()
		if err != nil {
			return errors.WithStack(err)
		}
		result.Value = int64(count)
		return nil
	})
	return result, err
}

func (s *Store) PackRefs(ctx context.Context, none *pb.None) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		return errors.WithStack(r.Storer.PackRefs())
	})
	return result, err
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

func (s *Store) getObject(uuid string) (*EncodedObject, bool) {
	return s.objectLRU.Get(uuid)
}
func (s *Store) putObject(obj *EncodedObject) {
	s.objectLRU.Put(obj)
}
