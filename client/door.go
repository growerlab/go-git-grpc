package client

import (
	"bufio"
	"bytes"
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

func (d *Door) RunGit(params *git.Context) error {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("ServeUploadPack panic: %+v", e)
		}
	}()

	runGit, err := d.client.RunGit(d.ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer runGit.CloseSend()

	if err = d.sendContextPack(runGit, params); err != nil {
		return err
	}

	return d.copy(runGit, params.In, params.Out)
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
			raw := scanner.Bytes()
			err = pipe.Send(&pb.Request{Raw: raw})
			if err != nil {
				log.Printf("read err: %+v\n", err)
				break
			}
		}
		log.Printf("scan is done.\n")
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
			if bytes.HasSuffix(resp.Raw, []byte("\r\nEOF")) {
				break
			}
			_, err = out.Write(resp.Raw)
			if err != nil {
				log.Printf("write: %+v\n", err)
				break
			}
		}
		log.Printf("recv is done.\n")
	}()

	wg.Wait()
	return
}

type clientStream interface {
	Send(*pb.Request) error
	Recv() (*pb.Response, error)
}

func (d *Door) sendContextPack(pack clientStream, params *git.Context) error {
	firstReq := &pb.Request{
		Path:     params.RepoPath,
		Env:      params.Env,
		GitBin:   params.GitBin,
		Args:     params.Args,
		Deadline: uint64(params.Deadline),
		Raw:      nil,
	}
	err := pack.Send(firstReq)
	return errors.WithStack(err)
}
