package external_svc

import (
	"alfred/modules/stack/v1/pb"
	"context"
)

func (s *stacksServer) GetStackBuild(ctx context.Context, in *pb.GetStackBuildRequest) (*pb.StackBuild, error) {
	return nil, nil
}
