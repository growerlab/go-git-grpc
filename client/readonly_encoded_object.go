package client

import (
	"io"

	"github.com/go-git/go-git/v5/plumbing"
)

var _ plumbing.EncodedObject = (*ReadonlyEncodedObject)(nil)

type ReadonlyEncodedObject struct {
	hash plumbing.Hash
	typ  plumbing.ObjectType
	size int64
}

func (t *ReadonlyEncodedObject) Hash() plumbing.Hash { return t.hash }

func (t *ReadonlyEncodedObject) Type() plumbing.ObjectType { return t.typ }

func (t *ReadonlyEncodedObject) SetType(plumbing.ObjectType) {}

func (t *ReadonlyEncodedObject) Size() int64 { return t.size }

func (t *ReadonlyEncodedObject) SetSize(int64) {}

func (t *ReadonlyEncodedObject) Reader() (io.ReadCloser, error) { return nil, nil }

func (t *ReadonlyEncodedObject) Writer() (io.WriteCloser, error) { return nil, nil }
