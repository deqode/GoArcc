package implementation

import (
	"alfred/modules/UserProfile/pb"
)

type UserServer struct {
}

func NewUserServer() pb.UserProfileServiceServer {
	return &UserServer{}
}
