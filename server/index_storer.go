package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

type IndexStorer struct {
}

func (s *IndexStorer) SetIndex(ctx context.Context, index *pb.Index) (*pb.None, error) {
	panic("implement me")
}

func (s *IndexStorer) GetIndex(ctx context.Context, none *pb.None) (*pb.Index, error) {
	panic("implement me")
}
