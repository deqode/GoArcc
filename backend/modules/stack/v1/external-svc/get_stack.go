package external_svc

import (
	"alfred/modules/stack/v1/pb"
	"context"
)

func (s *stacksServer) GetStack(ctx context.Context, in *pb.GetStackRequest) (*pb.Stack, error) {
	return nil, nil
}
