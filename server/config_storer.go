package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

type ConfigStorer struct {
}

func (s *ConfigStorer) GetConfig(ctx context.Context, none *pb.None) (*pb.Config, error) {
	panic("implement me")
}

func (s *ConfigStorer) SetConfig(ctx context.Context, c *pb.Config) (*pb.None, error) {
	panic("implement me")
}
