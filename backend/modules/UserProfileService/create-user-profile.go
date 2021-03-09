package UserProfileService

import (
	"alfred/logger"
	hello "alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService/pb"
	"context"
	"fmt"
	"os/exec"
)

func (server *UserProfileService) CreateUserProfile(ctx context.Context, request *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	var resp pb.UserProfile
	_, err := hello.LocalHelloWorld().HelloWorld(ctx, &hellopb.Hello{
		Message: "Hii",
	})
	if err != nil {
		panic(err)
	}
	//creating uuid
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Log.Debug("unable to generate uuid")
	}

	resp.Id = fmt.Sprintf("%s", out)
	resp.Name = request.UserProfile.Name
	resp.Email = request.UserProfile.Email
	resp.PhoneNumber = request.UserProfile.PhoneNumber
	resp.UserId = request.UserProfile.UserId
	t := server.db.Create(&resp)
	if t.Error != nil {
		return nil, t.Error
	}

	return &resp, nil

}
