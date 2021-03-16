package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

var _ pb.StorerServer = &Storer{}

type Storer struct {
	*pb.UnimplementedStorerServer
}

func (s *Storer) NewEncodedObject(ctx context.Context, none *pb.None) (*pb.EncodedObject, error) {
	panic("implement me")
}

func (s *Storer) SetEncodedObjectType(ctx context.Context, i *pb.Int8) (*pb.Int8, error) {
	panic("implement me")
}

func (s *Storer) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
	panic("implement me")
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
