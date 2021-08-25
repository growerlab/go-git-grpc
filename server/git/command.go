package git

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const DefaultTimeout = 60 * time.Minute // 推送和拉取的最大执行时间

const (
	UploadPackServiceName  = "git-upload-pack"
	ReceivePackServiceName = "git-receive-pack"
)

type Context struct {
	Env      map[string]string // 环境变量
	Rpc      string            // git upload or receive
	Args     []string          // args
	In       io.Reader         // input
	Out      io.Writer         // output
	RepoPath string            // repo dir

	Timeout time.Duration // 命令执行时间，单位秒
}

func Run(root string, params *Context) error {
	if params.Timeout <= 0 {
		params.Timeout = DefaultTimeout
	}

	// deadline
	cmdCtx, cancel := context.WithTimeout(context.Background(), params.Timeout)
	defer cancel()

	switch params.Rpc {
	case ReceivePackServiceName, UploadPackServiceName:
	default:
		return errors.Errorf("invalid rpc '%s'", params.Rpc)
	}

	cmd := exec.CommandContext(cmdCtx, params.Rpc, params.Args...)
	if len(params.Env) > 0 {
		systemEnvs := os.Environ()
		cmd.Env = make([]string, 0, len(params.Env)+len(systemEnvs))
		for k, v := range params.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = append(cmd.Env, os.Environ()...)
	}

	cmd.Dir = filepath.Join(root, params.RepoPath)
	if params.In != nil {
		cmd.Stdin = params.In
	}
	if params.Out != nil {
		cmd.Stdout = params.Out
	}
	cmd.Stderr = params.Out
	err := cmd.Run()
	return errors.WithStack(err)
}
