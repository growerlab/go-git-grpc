package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	gggrpc "github.com/growerlab/go-git-grpc"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	go func() {
		gitRoot := path.Join(os.Getenv("GOPATH"), "src/github.com/growerlab/go-git-grpc")
		err := gggrpc.NewServer(gitRoot, "localhost:8081")
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	clientCtx := context.Background()
	repoPath := "test/testrepo_bare"
	store, closeFn, err := gggrpc.NewClient(clientCtx, "localhost:8081", repoPath)
	if err != nil {
		panic(err)
	}
	defer closeFn.Close()

	repo, err := git.Open(store, nil)
	if err != nil {
		panic(err)
	}
	refs, err := repo.References()
	if err != nil {
		panic(err)
	}
	refs.ForEach(func(r *plumbing.Reference) error {
		fmt.Println(r.Name())
		return nil
	})
}
