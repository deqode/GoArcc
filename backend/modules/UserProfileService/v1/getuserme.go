package UserProfileService

import (
	"alfred/modules/UserProfileService/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *UserProfileService) GetUserMe(context.Context, *empty.Empty) (*pb.UserProfile, error) {
	panic("Implement me")
}