package git

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const DefaultTimeout = 60 * time.Minute // 推送和拉取的最大执行时间

type Context struct {
	Env      []string  // 环境变量 key=value
	GitBin   string    // git binary
	Args     []string  // git args
	In       io.Reader // git input
	Out      io.Writer // git output
	RepoPath string    // repo dir

	Deadline time.Duration // 命令执行时间，单位秒
}

func Run(root string, params *Context) error {
	if params.Deadline <= 0 {
		params.Deadline = DefaultTimeout
	}

	// deadline
	cmdCtx, cancel := context.WithTimeout(context.Background(), params.Deadline)
	defer cancel()

	cmd := exec.CommandContext(cmdCtx, params.GitBin, params.Args...)
	if len(params.Env) > 0 {
		systemEnvs := os.Environ()
		cmd.Env = make([]string, 0, len(params.Env)+len(systemEnvs))
		for _, v := range params.Env {
			cmd.Env = append(cmd.Env, v)
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
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return errors.WithStack(err)
}
