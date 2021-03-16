package server

import (
	"context"

	"github.com/growerlab/go-git-grpc/pb"
)

type ReferenceStorer struct {
}

func (s *ReferenceStorer) SetReference(ctx context.Context, reference *pb.Reference) (*pb.None, error) {
	panic("implement me")
}

func (s *ReferenceStorer) CheckAndSetReference(ctx context.Context, params *pb.SetReferenceParams) (*pb.None, error) {
	panic("implement me")
}

func (s *ReferenceStorer) GetReference(ctx context.Context, name *pb.ReferenceName) (*pb.Reference, error) {
	panic("implement me")
}

func (s *ReferenceStorer) GetReferences(ctx context.Context, none *pb.None) (*pb.References, error) {
	panic("implement me")
}

func (s *ReferenceStorer) RemoveReference(ctx context.Context, name *pb.ReferenceName) (*pb.None, error) {
	panic("implement me")
}

func (s *ReferenceStorer) CountLooseRefs(ctx context.Context, none *pb.None) (*pb.Int64, error) {
	panic("implement me")
}

func (s *ReferenceStorer) PackRefs(ctx context.Context, none *pb.None) (*pb.None, error) {
	panic("implement me")
}
