package internals

import (
	"alfred/modules/VCSConnection/v1/internals/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *vcsConnectionInternalServer) RevokeVCSToken(context.Context, *pb.RevokeVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
