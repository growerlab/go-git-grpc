package server

import (
	"time"

	"github.com/growerlab/go-git-grpc/common"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/growerlab/go-git-grpc/server/git"
	"github.com/pkg/errors"
)

// ServerCommand is used for a single server command execution.
type ServerCommand struct {
	repoPath   string
	uploadPack pb.Door_ServeUploadPackServer

	ctx *git.Context
}

// Start 协议：第一个请求仅传git相关的参数，不传数据
func (s *ServerCommand) Start() error {
	firstReq, err := s.uploadPack.Recv()
	if err != nil {
		return errors.WithStack(err)
	}

	s.ctx = &git.Context{
		Env:     common.ArrayToSet(firstReq.Env),
		Rpc:     firstReq.RPC,
		Args:    common.ArrayToArgs(firstReq.Args),
		In:      s,
		Out:     s,
		RepoDir: firstReq.Path,
		Timeout: time.Duration(firstReq.Timeout) * time.Second,
	}
	s.repoPath = firstReq.Path
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

// ServeUploadPack for git-upload-pack
func (d *Door) ServeUploadPack(pack pb.Door_ServeUploadPackServer) error {
	srvCmd := ServerCommand{uploadPack: pack}
	if err := srvCmd.Start(); err != nil {
		return err
	}

	return git.Run(srvCmd.ctx)
}

// ServeReceivePack for git-receive-pack
func (d *Door) ServeReceivePack(pack pb.Door_ServeReceivePackServer) error {
	srvCmd := ServerCommand{uploadPack: pack}
	if err := srvCmd.Start(); err != nil {
		return err
	}
	return git.Run(srvCmd.ctx)
}

func (d *Door) mustEmbedUnimplementedDoorServer() {}
