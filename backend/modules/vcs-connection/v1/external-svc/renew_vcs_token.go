package external_svc

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *vcsConnectionServer) RenewVCSToken(ctx context.Context, req *pb.RenewVCSTokenRequest) (*empty.Empty, error) {
	if err := s.ValidateUser(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}
