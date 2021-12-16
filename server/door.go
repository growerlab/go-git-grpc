package server

import (
	"log"
	"time"

	"github.com/growerlab/go-git-grpc/pb"
	"github.com/growerlab/go-git-grpc/server/git"
)

// ServerCommand is used for a single server command execution.
type ServerCommand struct {
	repoPath     string
	gitBinServer pb.Door_RunGitServer

	ctx *git.Context
}

// Start 协议：第一个请求仅传git相关的参数，不传数据
func (s *ServerCommand) Start() error {
	firstReq, err := s.gitBinServer.Recv()
	if err != nil {
		return err
	}

	s.ctx = &git.Context{
		Env:      firstReq.Env,
		GitBin:   firstReq.GitBin,
		Args:     firstReq.Args,
		In:       s,
		Out:      s,
		RepoPath: firstReq.Path,
		Deadline: time.Duration(firstReq.Deadline),
	}
	s.repoPath = firstReq.Path

	log.Println("---->")
	log.Println(s.ctx.String())

	return nil
}

func (s *ServerCommand) Read(p []byte) (n int, err error) {
	req, err := s.gitBinServer.Recv()
	if err != nil {
		if req == nil {
			return 0, err
		}
		return len(req.Raw), err
	}
	return copy(p, req.Raw), nil
}

func (s *ServerCommand) Write(p []byte) (n int, err error) {
	err = s.gitBinServer.Send(&pb.Response{Raw: p})
	return len(p), err
}

func (s *ServerCommand) Close() error {
	return nil
}

var _ pb.DoorServer = (*Door)(nil)

func NewDoor(root string) *Door {
	return &Door{
		root: root,
	}
}

type Door struct {
	*pb.UnimplementedDoorServer
	root string
}

// RunGit 执行git命令
func (d *Door) RunGit(pack pb.Door_RunGitServer) error {
	srvCmd := ServerCommand{gitBinServer: pack}
	if err := srvCmd.Start(); err != nil {
		return err
	}

	return git.Run(d.root, srvCmd.ctx, func(err error) {
		if err != nil {
			log.Printf("git.Run error: %v\n", err)
		}
		_ = pack.Send(&pb.Response{Raw: []byte("\r\nEOF")})
	})
}

func (d *Door) mustEmbedUnimplementedDoorServer() {}
