package server

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
)

func repo(root, path string, repoFn func(*git.Repository)) error {
	dir := filepath.Join(root, path)
	r, err := git.PlainOpen(dir)
	if err != nil {
		return errors.WithStack(err)
	}
	repoFn(r)
	return nil
}
