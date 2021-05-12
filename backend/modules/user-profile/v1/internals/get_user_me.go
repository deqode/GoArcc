package internals

import (
	"alfred/modules/user-profile/v1/internals/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *userProfileInternalServer) GetUserMe(context.Context, *empty.Empty) (*pb.UserProfile, error) {
	panic("Implement me")
}
