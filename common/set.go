package common

import (
	"github.com/growerlab/go-git-grpc/pb"
)

func ArrayToSet(values []*pb.KeyValue) map[string]string {
	result := map[string]string{}
	for _, v := range values {
		result[v.Key] = v.Value
	}
	return result
}

func ArrayToArgs(values []*pb.KeyValue) []string {
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.Key, v.Value)
	}
	return result
}
