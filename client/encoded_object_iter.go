package client

import (
	"context"
	"log"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

var _ storer.EncodedObjectIter = (*EncodedObjectIter)(nil)

type EncodedObjectIter struct {
	repoPath string
	ctx      context.Context
	client   pb.StorerClient

	objectType plumbing.ObjectType

	// 服务器端返回的iter none
	none *pb.None
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

	c.none, err = c.client.NewEncodedObjectIter(ctx, t)
	return c, err
}

func (c *EncodedObjectIter) Next() (plumbing.EncodedObject, error) {
	pbEncodedObject, err := c.client.EncodedObjectNext(c.ctx, c.none)
	if err == nil {
		return nil, err
	}

	return buildEncodedObjectFromPB(c.ctx, c.client, c.repoPath, pbEncodedObject), nil
}

func (c *EncodedObjectIter) ForEach(f func(plumbing.EncodedObject) error) error {
	eachClient, err := c.client.EncodedObjectForEach(c.ctx, c.none)
	if err != nil {
		return errors.WithStack(err)
	}

	forTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	for {
		select {
		case <-forTimeout.Done():
			return nil
		default:
			obj, err := eachClient.Recv()
			if err != nil {
				cancel()
				return err
			}
			err = f(buildEncodedObjectFromPB(c.ctx, c.client, c.repoPath, obj))
			if err != nil {
				return err
			}
		}
	}
}

func (c *EncodedObjectIter) Close() {
	_, err := c.client.EncodedObjectClose(c.ctx, c.none)
	if err != nil {
		log.Printf("call EncodedObjectClose was err: %+v\n", err)
	}
}
