package client

import (
	"context"
	"io"
	"log"

	"github.com/growerlab/go-git-grpc/common"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/growerlab/go-git-grpc/server/git"
	"github.com/pkg/errors"
	"github.com/reactivex/rxgo/v2"
)

type Door struct {
	ctx    context.Context //
	client pb.DoorClient   //
}

func NewDoor(ctx context.Context, pbClient pb.DoorClient) *Door {
	door := &Door{
		ctx:    ctx,
		client: pbClient,
	}
	return door
}

func (d *Door) ServeReceivePack(params *git.Context) error {
	receivePack, err := d.client.ServeReceivePack(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	receiver := rxgo.Start([]rxgo.Supplier{BytesReaderSupplier(params.In)}).ForEach(func(i interface{}) {
		if buf, ok := i.([]byte); ok {
			err = receivePack.Send(&pb.Request{Raw: buf})
			log.Printf("send msg to door server was err: %+v\n", err)
		}
	}, func(err error) {
		log.Printf("reader msg was err: %+v\n", err)
	}, func() {})

	sender := rxgo.Start([]rxgo.Supplier{ReceivePackSupplier(receivePack)}).ForEach(func(i interface{}) {
		if buf, ok := i.([]byte); ok {
			_, err = params.Out.Write(buf)
			log.Printf("write msg to Out was err: %+v\n", err)
		}
	}, func(err error) {
		log.Printf("recevie pack was err: %+v", err)
	}, func() {})

	<-receiver
	<-sender
	return nil
}

func (d *Door) ServeUploadPack(params *git.Context) error {
	uploadPack, err := d.client.ServeUploadPack(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *Door) sendInitPack(params *git.Context) {
	firstReq := &pb.Request{
		Path:    params.RepoPath,
		Env:     common.SetToPBSet(params.Env),
		RPC:     params.Rpc,
		Args:    params.Args,
		Timeout: uint64(params.Timeout),
		Raw:     nil,
	}
	err = receivePack.Send(firstReq)
	if err != nil {
		return errors.WithStack(err)
	}
}

func ReceivePackSupplier(receivePack pb.Door_ServeReceivePackClient) rxgo.Supplier {
	return func(ctx context.Context) rxgo.Item {
		pack, err := receivePack.Recv()
		if err != nil {
			return rxgo.Error(err)
		}
		return rxgo.Of(pack.Raw)
	}
}

func BytesReaderSupplier(reader io.Reader) rxgo.Supplier {
	return func(ctx context.Context) rxgo.Item {
		// 也许未来某个时刻需要引入 bytes pool
		var buf = make([]byte, 1024)
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			return rxgo.Error(err)
		}
		return rxgo.Of(buf)
	}
}
