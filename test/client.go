package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/growerlab/go-git-grpc/client"
	"github.com/growerlab/go-git-grpc/server"
)

func main() {

	go func() {
		gitRoot := path.Join(os.Getenv("GOPATH"), "src/github.com/growerlab/go-git-grpc")
		err := server.New(gitRoot, "localhost:8081")
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	clientCtx := context.Background()
	repoPath := "test/testrepo_bare"
	store, closeFn, err := client.New(clientCtx, "localhost:8081", repoPath)
	if err != nil {
		panic(err)
	}
	defer closeFn.Close()

	repo, err := git.Open(store, nil)
	if err != nil {
		panic(err)
	}
	bn, err := repo.Branch("master")
	if err != nil {
		panic(err)
	}
	fmt.Println(bn)
}
