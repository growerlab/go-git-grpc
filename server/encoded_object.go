package server

import (
	"context"
	"io"

	"github.com/go-git/go-git/v5/plumbing"
)

var _ EncodedObjectExt = (*EncodedObject)(nil)

type EncodedObjectExt interface {
	UUID() string
	plumbing.EncodedObject
}

type EncodedObject struct {
	ctx context.Context

	uuid     string
	repoPath string
	obj      plumbing.EncodedObject
}

func (o *EncodedObject) UUID() string { return o.uuid }

func (o *EncodedObject) Hash() plumbing.Hash { return o.obj.Hash() }

func (o *EncodedObject) Type() plumbing.ObjectType { return o.obj.Type() }

func (o *EncodedObject) SetType(t plumbing.ObjectType) { o.obj.SetType(t) }

func (o *EncodedObject) Size() int64 { return o.obj.Size() }

func (o *EncodedObject) SetSize(s int64) { o.obj.SetSize(s) }

func (o *EncodedObject) Reader() (io.ReadCloser, error) {
	return o.obj.Reader()
}

func (o *EncodedObject) Writer() (io.WriteCloser, error) {
	return o.obj.Writer()
}
