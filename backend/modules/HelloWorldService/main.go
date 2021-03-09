package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
	usrprofile "alfred/modules/UserProfileService/pb"
)

type HelloService struct {
	userProfileServiceClient usrprofile.UserProfileServiceClient
}

//Service Implementation
func NewHelloService(userProfileServiceClient usrprofile.UserProfileServiceClient) pb.HelloWorldServiceServer {
	return HelloService{
		userProfileServiceClient: userProfileServiceClient,
	}
}

func LocalHelloWorld() pb.HelloWorldServiceServer {
	return HelloService{}
}
