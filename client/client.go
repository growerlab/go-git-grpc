package client

import (
	"io"
	"log"

	"github.com/growerlab/go-git-grpc/pb"
	"google.golang.org/grpc"
)

func New(address string) (*Store, io.Closer) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewStorerClient(conn)
	s := &Store{
		grpcConn:    conn,
		storeClient: c,
	}
	return s, s
}
