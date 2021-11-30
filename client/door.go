package client

import (
	"bufio"
	"context"
	"io"
	"log"
	"sync"

	"github.com/growerlab/go-git-grpc/pb"
	"github.com/growerlab/go-git-grpc/server/git"
	"github.com/pkg/errors"
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
	defer func() {
		if e := recover(); e != nil {
			log.Printf("ServeReceivePack panic: %+v", e)
		}
	}()

	receivePack, err := d.client.ServeReceivePack(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = d.sendContextPack(receivePack, params); err != nil {
		return err
	}
	return d.copy(receivePack, params.In, params.Out)
}

func (d *Door) ServeUploadPack(params *git.Context) error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("ServeUploadPack panic: %+v", e)
		}
	}()

	uploadPack, err := d.client.ServeUploadPack(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = d.sendContextPack(uploadPack, params); err != nil {
		return err
	}

	return d.copy(uploadPack, params.In, params.Out)
}

func (d *Door) copy(pipe clientStream, in io.Reader, out io.Writer) (err error) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if in == nil {
			return
		}
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			err = pipe.Send(&pb.Request{Raw: scanner.Bytes()})
			if err != nil {
				log.Printf("read err: %+v\n", err)
				break
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("scan err: %+v\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var resp *pb.Response
			resp, err = pipe.Recv()
			if err != nil {
				log.Printf("receive: %+v\n", err)
				break
			}
			_, err = out.Write(resp.Raw)
			if err != nil {
				log.Printf("write: %+v\n", err)
				break
			}
		}
	}()

	wg.Wait()
	return err
}

type clientStream interface {
	Send(*pb.Request) error
	Recv() (*pb.Response, error)
}

func (d *Door) sendContextPack(pack clientStream, params *git.Context) error {
	firstReq := &pb.Request{
		Path:    params.RepoPath,
		Env:     params.Env,
		RPC:     params.Rpc,
		Args:    params.Args,
		Timeout: uint64(params.Timeout),
		Raw:     nil,
	}
	err := pack.Send(firstReq)
	return errors.WithStack(err)
}
