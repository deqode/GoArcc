package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
	usrprofile "alfred/modules/UserProfileService/pb"
)

type HelloService struct {
	usrCli usrprofile.UserProfileServiceClient
}

//Service Implementation
func NewHelloService(usrCli usrprofile.UserProfileServiceClient) pb.HelloWorldServiceServer {
	return HelloService{
		usrCli: usrCli,
	}
}
