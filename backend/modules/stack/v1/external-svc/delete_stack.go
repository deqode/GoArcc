package external_svc

import (
	"alfred/modules/stack/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *stacksServer) DeleteStack(ctx context.Context, in *pb.DeleteStackRequest) (*empty.Empty, error) {
	return nil, nil
}
