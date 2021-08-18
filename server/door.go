package server

import (
	"context"
	"fmt"
	"io"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/server"
	"github.com/go-git/go-git/v5/utils/ioutil"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/reactivex/rxgo/v2"
)

// ServerCommand is used for a single server command execution.
type ServerCommand struct {
	Stderr io.Writer
	Stdout io.WriteCloser
	Stdin  io.Reader
}

var _ pb.DoorServer = (*Door)(nil)

type Door struct {
}

func (d *Door) ServeUploadPack(pack pb.Door_ServeUploadPackServer) error {
	ep, err := transport.NewEndpoint(path)
	if err != nil {
		return err
	}

	firstPack := pack.Recv()

	rxgo.FromEventSource()

	reader, writer := io.Pipe()

	srvCmd := ServerCommand{
		Stderr: nil,
		Stdout: nil,
		Stdin:  nil,
	}

	// 在服务端不检测授权，由客户端做
	s, err := server.DefaultServer.NewUploadPackSession(ep, nil)
	if err != nil {
		return fmt.Errorf("error creating session: %s", err)
	}

	return d.serveUploadPack(srvCmd, s)
}

func (d *Door) ServeReceivePack(pack pb.Door_ServeReceivePackServer) error {

}

func (d *Door) serveUploadPack(cmd ServerCommand, s transport.UploadPackSession) (err error) {
	ioutil.CheckClose(cmd.Stdout, &err)

	ar, err := s.AdvertisedReferences()
	if err != nil {
		return err
	}

	if err := ar.Encode(cmd.Stdout); err != nil {
		return err
	}

	req := packp.NewUploadPackRequest()
	if err := req.Decode(cmd.Stdin); err != nil {
		return err
	}

	var resp *packp.UploadPackResponse
	resp, err = s.UploadPack(context.TODO(), req)
	if err != nil {
		return err
	}

	return resp.Encode(cmd.Stdout)
}

func (d *Door) serveReceivePack(cmd ServerCommand, s transport.ReceivePackSession) error {
	ar, err := s.AdvertisedReferences()
	if err != nil {
		return fmt.Errorf("internal error in advertised references: %s", err)
	}

	if err := ar.Encode(cmd.Stdout); err != nil {
		return fmt.Errorf("error in advertised references encoding: %s", err)
	}

	req := packp.NewReferenceUpdateRequest()
	if err := req.Decode(cmd.Stdin); err != nil {
		return fmt.Errorf("error decoding: %s", err)
	}

	rs, err := s.ReceivePack(context.TODO(), req)
	if rs != nil {
		if err := rs.Encode(cmd.Stdout); err != nil {
			return fmt.Errorf("error in encoding report status %s", err)
		}
	}

	if err != nil {
		return fmt.Errorf("error in receive pack: %s", err)
	}

	return nil
}

func (d *Door) mustEmbedUnimplementedDoorServer() {}
