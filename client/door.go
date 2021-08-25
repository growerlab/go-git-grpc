package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/growerlab/go-git-grpc/pb"
)

type Door struct {
	repoPath string          // 仓库路径
	ctx      context.Context //
	client   pb.DoorClient   //
}

func NewDoor(ctx context.Context, repoPath string, pbClient pb.DoorClient) *Door {
	door := &Door{
		repoPath: repoPath,
		ctx:      ctx,
		client:   pbClient,
	}
	return door
}

func (d *Door) ServeReceivePack() error {
	receive, err := d.client.ServeReceivePack(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	req := &pb.Request{
		Path:    "",
		Env:     nil,
		RPC:     "",
		Args:    nil,
		Timeout: 0,
		Raw:     nil,
	}
	receive.Send(req)

	return nil
}

func (d *Door) ServeUploadPack() error {
	d.client.ServeUploadPack()
	return nil
}
