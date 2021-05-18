package user_profile

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *userProfileInServer) GetUserMe(context.Context, *empty.Empty) (*pb.UserProfile, error) {
	panic("Implement me")
}
