package external_svc

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *vcsConnectionServer) RevokeVCSToken(ctx context.Context, req *pb.RevokeVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
