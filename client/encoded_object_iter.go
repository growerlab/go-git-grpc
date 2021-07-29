package client

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

var _ storer.EncodedObjectIter = (*CustomEncodedObjectIter)(nil)

type CustomEncodedObjectIter struct {
}

func (c *CustomEncodedObjectIter) Next() (plumbing.EncodedObject, error) {
	panic("implement me")
}

func (c *CustomEncodedObjectIter) ForEach(f func(plumbing.EncodedObject) error) error {
	panic("implement me")
}

func (c *CustomEncodedObjectIter) Close() {
	panic("implement me")
}
