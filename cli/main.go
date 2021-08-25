package main

import (
	"github.com/go-git/go-git/v5"
	gggrpc "github.com/growerlab/go-git-grpc"
)

func main() {
	git.Open(nil, nil)

	gggrpc.NewServer("./", ":8081")

	gggrpc.NewStoreClient(nil, "", "")
}
