package client

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/growerlab/go-git-grpc/pb"
)

var _ storer.EncodedObjectIter = (*EncodedObjectIter)(nil)

type EncodedObjectIter struct {
	repoPath string
	ctx      context.Context
	client   pb.StorerClient

	objectType plumbing.ObjectType

	// 服务器端返回的iter uuid
	uuid *pb.None
}

func NewEncodedObjectIter(ctx context.Context, client pb.StorerClient, repoPath string, objectType plumbing.ObjectType) (*EncodedObjectIter, error) {
	var err error
	c := &EncodedObjectIter{
		ctx:        ctx,
		repoPath:   repoPath,
		client:     client,
		objectType: objectType,
	}

	t := &pb.ObjectType{Type: objectType.String()}

	c.uuid, err = c.client.NewEncodedObjectIter(ctx, t)
	return c, err
}

func (c *EncodedObjectIter) Next() (plumbing.EncodedObject, error) {
	obj, err := c.client.EncodedObjectNext(c.ctx, c.uuid)
	if err != nil {
		return nil, err
	}

}

func (c *EncodedObjectIter) ForEach(f func(plumbing.EncodedObject) error) error {
	panic("implement me")
}

func (c *EncodedObjectIter) Close() {
	panic("implement me")
}
