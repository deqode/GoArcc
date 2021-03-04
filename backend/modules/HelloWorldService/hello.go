package HelloWorldService

import (
	"alfred/logger"
	"alfred/modules/HelloWorldService/pb"
	usrprofilePb "alfred/modules/UserProfileService/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h HelloService) HelloWorld(ctx context.Context, hello *pb.Hello) (*empty.Empty, error) {
	logger.Log.Info("Hello")
	_, err := h.userProfileServiceClient.CreateUserProfile(ctx, &usrprofilePb.CreateUserProfileRequest{UserProfile: &usrprofilePb.UserProfile{
		Name: "Tanuj sir",
	}})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
