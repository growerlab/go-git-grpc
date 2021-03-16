package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

type ShallowStorer struct {
}

func (s *ShallowStorer) SetShallow(ctx context.Context, hashs *pb.Hashs) (*pb.None, error) {
	panic("implement me")
}

func (s *ShallowStorer) Shallow(ctx context.Context, none *pb.None) (*pb.Hashs, error) {
	panic("implement me")
}
