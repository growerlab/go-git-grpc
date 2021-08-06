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

	//
	nextClient pb.Storer_EncodedObjectNextClient
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
	var err error
	if c.nextClient == nil {
		c.nextClient, err = c.client.EncodedObjectNext(c.ctx, c.none)
		if err != nil {
			return nil, err
		}
	}

	pbEncodedObject, err := c.nextClient.Recv()
	if err == nil {
		return nil, err
	}

	return &EncodedObject{
		ctx:            c.ctx,
		client:         c.client,
		repoPath:       c.repoPath,
		uuid:           pbEncodedObject.UUID,
		readonlyObject: buildReadonlyEncodedObject(pbEncodedObject),
	}, nil
}

func (c *EncodedObjectIter) ForEach(f func(plumbing.EncodedObject) error) error {
	eachClient, err := c.client.EncodedObjectForEach(c.ctx, c.none)
	if err != nil {
		return errors.WithStack(err)
	}

	forTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
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
			err = f(buildReadonlyEncodedObject(obj))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *EncodedObjectIter) Close() {
	_, err := c.client.EncodedObjectClose(c.ctx, c.none)
	log.Println("Call EncodedObjectClose was err: %+v\n", err)
}
