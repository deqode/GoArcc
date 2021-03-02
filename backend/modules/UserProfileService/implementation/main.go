package implementation

import (
	"alfred/modules/UserProfileService/pb"
)

type UserProfileService struct {
}

func NewUserProfileService() pb.UserProfileServiceServer {
	return &UserProfileService{}
}
