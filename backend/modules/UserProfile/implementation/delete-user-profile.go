package implementation

import (
	"alfred/modules/UserProfile/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (u *UserServer) DeleteUserProfile(ctx context.Context, request *pb.DeleteUserProfileRequest) (*empty.Empty, error) {
	panic("implement me")
}
