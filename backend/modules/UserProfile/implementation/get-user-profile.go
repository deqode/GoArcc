package implementation

import (
	"alfred/modules/UserProfile/pb"
	"context"
)

func (u *UserServer) GetUserProfile(ctx context.Context, request *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}
