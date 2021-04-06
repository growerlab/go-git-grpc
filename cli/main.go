package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/growerlab/go-git-grpc/client"
	"github.com/growerlab/go-git-grpc/server"
)

func main() {
	git.Open(nil, nil)

	server.New("./", ":8081")

	client.New(nil, "", "")
}
