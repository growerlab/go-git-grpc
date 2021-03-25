package server

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
)

func buildRefToPbRef(ref *plumbing.Reference) *pb.Reference {
	result := &pb.Reference{
		T:      []byte{byte(ref.Type())},
		N:      string(ref.Name()),
		H:      ref.Hash().String(),
		Target: string(ref.Target()),
	}
	return result
}
