package external_svc

import (
	"alfred/modules/stack/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *stacksServer) DeleteStackBuild(ctx context.Context, in *pb.DeleteStackBuildRequest) (*empty.Empty, error) {
	return nil, nil
}
