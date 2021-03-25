package server

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing"

	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5"
	"github.com/growerlab/go-git-grpc/pb"
)

var _ pb.StorerServer = &Storer{}

type Storer struct {
	*pb.UnimplementedStorerServer

	root string // 仓库根目录
}

func (s *Storer) NewEncodedObject(ctx context.Context, none *pb.None) (*pb.EncodedObject, error) {
	var result *pb.EncodedObject
	err := repo(s.root, none.RepoPath, func(r *git.Repository) {
		obj := r.Storer.NewEncodedObject()
		result = new(pb.EncodedObject)
		result.Type = obj.Type().String()
		result.Size = obj.Size()
		result.Hash = obj.Hash().String()
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) SetEncodedObjectType(ctx context.Context, i *pb.Int8) (*pb.None, error) {
	var result = &pb.None{RepoPath: i.RepoPath}
	var encodedObjectType plumbing.ObjectType

	if len(i.Value) > 0 {
		encodedObjectType = plumbing.ObjectType(i.Value[0])
	}

	err := repo(s.root, i.RepoPath, func(r *git.Repository) {
		r.Storer.NewEncodedObject().SetType(encodedObjectType)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
	var result = &pb.None{RepoPath: i.RepoPath}

	err := repo(s.root, i.RepoPath, func(r *git.Repository) {
		r.Storer.NewEncodedObject().SetSize(i.Value)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) EncodedObjectReader(none *pb.None, server pb.Storer_EncodedObjectReaderServer) error {
	panic("implement me")
}

func (s *Storer) EncodedObjectWriter(server pb.Storer_EncodedObjectWriterServer) error {
	panic("implement me")
}

func (s *Storer) SetReference(ctx context.Context, reference *pb.Reference) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) CheckAndSetReference(ctx context.Context, params *pb.SetReferenceParams) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) GetReference(ctx context.Context, name *pb.ReferenceName) (*pb.Reference, error) {
	panic("implement me")
}

func (s *Storer) GetReferences(ctx context.Context, none *pb.None) (*pb.References, error) {
	panic("implement me")
}

func (s *Storer) RemoveReference(ctx context.Context, name *pb.ReferenceName) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) CountLooseRefs(ctx context.Context, none *pb.None) (*pb.Int64, error) {
	panic("implement me")
}

func (s *Storer) PackRefs(ctx context.Context, none *pb.None) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) SetShallow(ctx context.Context, hashs *pb.Hashs) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) Shallow(ctx context.Context, none *pb.None) (*pb.Hashs, error) {
	panic("implement me")
}

func (s *Storer) SetIndex(ctx context.Context, index *pb.Index) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) GetIndex(ctx context.Context, none *pb.None) (*pb.Index, error) {
	panic("implement me")
}

func (s *Storer) GetConfig(ctx context.Context, none *pb.None) (*pb.Config, error) {
	panic("implement me")
}

func (s *Storer) SetConfig(ctx context.Context, c *pb.Config) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) Modules(ctx context.Context, none *pb.None) (*pb.ModuleNames, error) {
	panic("implement me")
}

func (s *Storer) mustEmbedUnimplementedStorerServer() {

}
