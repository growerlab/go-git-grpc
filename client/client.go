package client

import (
	"context"
	"io"

	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func New(ctx context.Context, grpcServerAddr string, repoPath string) (*Store, io.Closer, error) {
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
	s := &Store{
		repoPath: repoPath,
		grpcConn: conn,
		client:   c,
		ctx:      ctx,
	}
	return s, s, nil
}
