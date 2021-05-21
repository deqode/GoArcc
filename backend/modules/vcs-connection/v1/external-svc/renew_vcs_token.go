package external_svc

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *vcsConnectionServer) RenewVCSToken(context.Context, *pb.RenewVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
