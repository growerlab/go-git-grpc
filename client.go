package gggrpc

import (
	"context"
	"io"

	"github.com/growerlab/go-git-grpc/client"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, grpcServerAddr string, repoPath string) (*client.Store, io.Closer, error) {
	conn, err := grpc.DialContext(ctx,
		grpcServerAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		err := errors.WithStack(err)
		return nil, nil, err
	}

	c := pb.NewStorerClient(conn)
	s := client.NewStore(ctx, repoPath, conn, c)
	return s, s, nil
}
