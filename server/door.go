package server

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/server"
	"github.com/go-git/go-git/v5/utils/ioutil"
	"github.com/growerlab/go-git-grpc/pb"
)

// ServerCommand is used for a single server command execution.
type ServerCommand struct {
	Stderr io.Writer
	Stdout io.WriteCloser
	Stdin  io.Reader

	repoPath   string
	uploadPack pb.Door_ServeUploadPackServer
}

// Start 协议：第一个请求仅包含仓库路径，不传数据
func (s *ServerCommand) Start() error {
	firstReq, err := s.uploadPack.Recv()
	if err != nil {
		return errors.WithStack(err)
	}
	s.repoPath = firstReq.Path
	s.Stdout = s
	s.Stderr = s
	s.Stdin = s

	return nil
}

func (s *ServerCommand) Read(p []byte) (n int, err error) {
	req, err := s.uploadPack.Recv()
	if err != nil {
		return len(req.Raw), errors.WithStack(err)
	}
	return copy(p, req.Raw), nil
}

func (s *ServerCommand) Write(p []byte) (n int, err error) {
	err = s.uploadPack.Send(&pb.Response{Raw: p})
	return len(p), errors.WithStack(err)
}

func (s *ServerCommand) Close() error {
	return nil
}

var _ pb.DoorServer = (*Door)(nil)

type Door struct {
}

func (d *Door) ServeUploadPack(pack pb.Door_ServeUploadPackServer) error {
	srvCmd := ServerCommand{uploadPack: pack}
	if err := srvCmd.Start(); err != nil {
		return err
	}

	ep, err := transport.NewEndpoint(srvCmd.repoPath)
	if err != nil {
		return err
	}

	// 在服务端不检测授权，由客户端做
	s, err := server.DefaultServer.NewUploadPackSession(ep, nil)
	if err != nil {
		return fmt.Errorf("error creating session: %s", err)
	}

	return d.serveUploadPack(srvCmd, s)
}

func (d *Door) ServeReceivePack(pack pb.Door_ServeReceivePackServer) error {
	srvCmd := ServerCommand{uploadPack: pack}
	if err := srvCmd.Start(); err != nil {
		return err
	}

	ep, err := transport.NewEndpoint(srvCmd.repoPath)
	if err != nil {
		return err
	}

	s, err := server.DefaultServer.NewReceivePackSession(ep, nil)
	if err != nil {
		return fmt.Errorf("error creating session: %s", err)
	}

	return d.serveReceivePack(srvCmd, s)
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
