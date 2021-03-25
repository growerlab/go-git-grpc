package server

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
)

// TODO 未来可能需要加入LRU缓存针对打开的仓库对象handle
func repo(root, path string, repoFn func(*git.Repository) error) error {
	dir := filepath.Join(root, path)
	r, err := git.PlainOpen(dir)
	if err != nil {
		return errors.WithStack(err)
	}
	return repoFn(r)
}
