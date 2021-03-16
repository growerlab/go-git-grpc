package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

type EncodedObjectStorer struct {
}

func (s *EncodedObjectStorer) NewEncodedObject(ctx context.Context, none *pb.None) (*pb.EncodedObject, error) {
	panic("implement me")
}

func (s *EncodedObjectStorer) SetEncodedObjectType(ctx context.Context, i *pb.Int8) (*pb.Int8, error) {
	panic("implement me")
}

func (s *EncodedObjectStorer) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
	panic("implement me")
}

func (s *EncodedObjectStorer) EncodedObjectReader(none *pb.None, server pb.Storer_EncodedObjectReaderServer) error {
	panic("implement me")
}

func (s *EncodedObjectStorer) EncodedObjectWriter(server pb.Storer_EncodedObjectWriterServer) error {
	panic("implement me")
}
