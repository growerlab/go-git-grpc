package client

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/growerlab/go-git-grpc/pb"
)

var _ storer.EncodedObjectIter = (*EncodedObjectIterClient)(nil)

type EncodedObjectIterClient struct {
	repoPath string
	ctx      context.Context
	client   pb.StorerClient
}

func NewEncodedObjectIterClient(ctx context.Context, repoPath string, client pb.StorerClient) *EncodedObjectIterClient {
	return &EncodedObjectIterClient{
		ctx:      ctx,
		repoPath: repoPath,
		client:   client,
	}
}

func (c *EncodedObjectIterClient) Next() (plumbing.EncodedObject, error) {
	panic("implement me")
}

func (c *EncodedObjectIterClient) ForEach(f func(plumbing.EncodedObject) error) error {
	panic("implement me")
}

func (c *EncodedObjectIterClient) Close() {
	panic("implement me")
}
