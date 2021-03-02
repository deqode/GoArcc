package implementation

import (
	"alfred/modules/UserProfileService/pb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
)

func (u *UserProfileService) DeleteUserProfile(ctx context.Context, request *pb.DeleteUserProfileRequest) (*empty.Empty, error) {

	fmt.Print(u.conf.GRPCPort)
	panic("implement me")
}
