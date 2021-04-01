package server

import (
	"bytes"
	"context"
	"io"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
)

func NewEncodedObject() plumbing.EncodedObject {
	return &MemoryObject{}
}

var _ plumbing.EncodedObject = (*MemoryObject)(nil)

// MemoryObject on memory Object implementation
type MemoryObject struct {
	ctx    context.Context
	server pb.StorerServer

	t    plumbing.ObjectType
	h    plumbing.Hash
	cont []byte
	sz   int64
}

// Hash returns the object Hash, the hash is calculated on-the-fly the first
// time it's called, in all subsequent calls the same Hash is returned even
// if the type or the content have changed. The Hash is only generated if the
// size of the content is exactly the object size.
func (o *MemoryObject) Hash() plumbing.Hash {
	if o.h == plumbing.ZeroHash && int64(len(o.cont)) == o.sz {
		o.h = plumbing.ComputeHash(o.t, o.cont)
	}

	return o.h
}

// Type return the ObjectType
func (o *MemoryObject) Type() plumbing.ObjectType { return o.t }

// SetType sets the ObjectType
func (o *MemoryObject) SetType(t plumbing.ObjectType) { o.t = t }

// Size return the size of the object
func (o *MemoryObject) Size() int64 { return o.sz }

// SetSize set the object size, a content of the given size should be written
// afterwards
func (o *MemoryObject) SetSize(s int64) { o.sz = s }

// Reader returns an io.ReadCloser used to read the object's content.
//
// For a MemoryObject, this reader is seekable.
func (o *MemoryObject) Reader() (io.ReadCloser, error) {
	return nopCloser{bytes.NewReader(o.cont)}, nil
}

// Writer returns a ObjectWriter used to write the object's content.
func (o *MemoryObject) Writer() (io.WriteCloser, error) {
	return o, nil
}

func (o *MemoryObject) Write(p []byte) (n int, err error) {
	o.cont = append(o.cont, p...)
	o.sz = int64(len(o.cont))

	return len(p), nil
}

// Close releases any resources consumed by the object when it is acting as a
// ObjectWriter.
func (o *MemoryObject) Close() error { return nil }

// nopCloser exposes the extra methods of bytes.Reader while nopping Close().
//
// This allows clients to attempt seeking in a cached Blob's Reader.
type nopCloser struct {
	*bytes.Reader
}

// Close does nothing.
func (nc nopCloser) Close() error { return nil }
