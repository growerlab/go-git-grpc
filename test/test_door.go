package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	gggrpc "github.com/growerlab/go-git-grpc"
	"github.com/growerlab/go-git-grpc/client"
	"github.com/growerlab/go-git-grpc/server/git"
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
	in := io.NopCloser(strings.NewReader(""))
	out := strings.Builder{}

	cmd := &git.Context{
		Env:      []string{""},
		Rpc:      "git-upload-pack",
		Args:     []string{"--advertise-refs", "."},
		RepoPath: repoPath,
		In:       in,
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
