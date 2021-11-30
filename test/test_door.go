package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/growerlab/go-git-grpc/server/git"

	"github.com/growerlab/go-git-grpc/client"

	gggrpc "github.com/growerlab/go-git-grpc"
)

// 测试 git-upload-pack git-receive-pack 操作

var repoPath = "testrepo_bare"

func initServer() {
	gitRoot := os.Getenv("GO_GIT_GRPC_TEST_DIR")

	go func() {
		err := gggrpc.NewServer(gitRoot, "localhost:8081")
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)
}

func main() {
	initServer()

	door, closeFn, err := gggrpc.NewDoorClient(context.Background(), "localhost:8081")
	if err != nil {
		panic(err)
	}
	defer closeFn.Close()

	if err := testUploadCommand(door); err != nil {
		panic(err)
	}
	if err := testPush(); err != nil {
		panic(err)
	}
}

func testUploadCommand(door *client.Door) error {
	// 测试 git-upload-pack
	// in := bytes.Buffer{}
	out := bytes.Buffer{}

	cmd := &git.Context{
		Env:      []string{""},
		Rpc:      "git-upload-pack",
		Args:     []string{"--advertise-refs", "."},
		RepoPath: repoPath,
		In:       nil,
		Out:      &out,
	}

	if err := door.ServeUploadPack(cmd); err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}

func testPush() error {

	return nil
}

func testPull() error {
	return nil
}
