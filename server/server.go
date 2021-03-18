package server

import (
	"net"

	"github.com/growerlab/go-git-grpc/pb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func New(root, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store := &Storer{
		root: root,
	}

	pb.RegisterStorerServer(s, store)
	if err := s.Serve(lis); err != nil {
		return errors.Errorf("failed to serve: %v", err)
	}
	return nil
}
