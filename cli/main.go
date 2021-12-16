package main

import (
	"log"
	"os"
	"path/filepath"

	gggrpc "github.com/growerlab/go-git-grpc"
)

func main() {
	gitRoot := filepath.Join(os.Getenv("GOPATH"), "src", "github.com/growerlab/mensa/test/repos")

	log.Println("go-git-grpc running...")
	log.Println("git root:", gitRoot)

	err := gggrpc.NewServer(gitRoot, ":9001")
	if err != nil {
		panic(err)
	}
}
