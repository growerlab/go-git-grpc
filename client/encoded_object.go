package client

import (
	"context"
	"io"
	"log"

	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
)

var _ plumbing.EncodedObject = (*EncodedObject)(nil)

type EncodedObject struct {
	ctx      context.Context
	repoPath string
	client   pb.StorerClient

	obj *pb.EncodedObject
}

func (e *EncodedObject) Init() error {
	var err error
	var params = &pb.None{RepoPath: e.repoPath}

	e.obj, err = e.client.NewEncodedObject(e.ctx, params)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (e *EncodedObject) Hash() plumbing.Hash {
	return plumbing.NewHash(e.obj.Hash)
}

func (e *EncodedObject) Type() plumbing.ObjectType {
	ty, err := plumbing.ParseObjectType(e.obj.Type)
	if err != nil {
		log.Printf("Get encoded object type '%s' was err: %+v\n", e.obj.Type, err)
	}
	return ty
}

func (e *EncodedObject) SetType(objectType plumbing.ObjectType) {
	var params = &pb.Int8{
		RepoPath: e.repoPath,
		Value:    []byte{byte(objectType)},
	}
	_, err := e.client.SetEncodedObjectType(e.ctx, params)
	if err != nil {
		log.Printf("Set encoded object type '%s' was err: %+v\n", objectType.String(), err)
	}
}

func (e *EncodedObject) Size() int64 {
	return e.obj.Size
}

func (e *EncodedObject) SetSize(i int64) {
	var params = &pb.Int64{
		RepoPath: e.repoPath,
		Value:    i,
	}
	_, err := e.client.SetEncodedObjectSetSize(e.ctx, params)
	if err != nil {
		log.Printf("Set encoded object size was err: %+v\n", err)
	}
}

func (e *EncodedObject) Reader() (io.ReadCloser, error) {
	var params = &pb.None{RepoPath: e.repoPath}
	reader, err := e.client.EncodedObjectReader(e.ctx, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &EncodedObjectReadWriter{ctx: e.ctx, reader: reader}, nil
}

func (e *EncodedObject) Writer() (io.WriteCloser, error) {
	// var params = &pb.None{RepoPath: e.repoPath}
	writer, err := e.client.EncodedObjectWriter(e.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &EncodedObjectReadWriter{ctx: e.ctx, writer: writer}, nil
}

var _ io.ReadWriteCloser = (*EncodedObjectReadWriter)(nil)

type EncodedObjectReadWriter struct {
	ctx    context.Context
	reader pb.Storer_EncodedObjectReaderClient
	writer pb.Storer_EncodedObjectWriterClient
}

func (e *EncodedObjectReadWriter) Read(p []byte) (n int, err error) {
	if len(p) < 512 {
		return 0, errors.New("'p' minimum length 512")
	}
	b, err := e.reader.Recv()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	n = copy(p, b.Value)
	return
}

func (e *EncodedObjectReadWriter) Write(p []byte) (n int, err error) {

	panic("implement me")
}

func (e *EncodedObjectReadWriter) Close() error {
	if e.reader != nil {
		e.reader.CloseSend()
	}
}
