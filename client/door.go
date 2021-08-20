package client

import (
	"github.com/growerlab/go-git-grpc/pb"
)

type Door struct {
	addr string // 客户端的接口，支持git客户端的推拉仓库，通常通过 http、ssh

	client pb.DoorClient
}

func NewDoor(addr string) *Door {
	door := &Door{addr: addr}
	if err := door.start(); err != nil {
		panic(err)
	}
	return door
}

func (d *Door) start() error {
	d.client.ServeReceivePack()
	return nil
}
