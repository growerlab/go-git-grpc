package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

var _ pb.StorerServer = (*Storer)(nil)

type Storer struct {
	pb.UnimplementedStorerServer
	*EncodedObjectStorer
	*ReferenceStorer
	*ShallowStorer
	*IndexStorer
	*ConfigStorer
}

func (s *Storer) Modules(ctx context.Context, none *pb.None) (*pb.ModuleNames, error) {
	panic("implement me")
}

func (s *Storer) mustEmbedUnimplementedStorerServer() {

}
