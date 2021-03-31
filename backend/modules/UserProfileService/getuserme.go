package UserProfileService

import (
	"alfred/modules/UserProfileService/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (server *UserProfileService) GetUserMe(context.Context, *empty.Empty) (*pb.UserProfile, error) {
	panic("Implement me")
}
