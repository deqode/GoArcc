package UserProfileService

import (
	"alfred/logger"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService/pb"
	"context"
	"fmt"
	"os/exec"
)

func (server *UserProfileService) CreateUserProfile(ctx context.Context, request *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	var resp pb.UserProfile

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
	_, err = server.helloCli.HelloWorld(ctx, &hellopb.Hello{
		Message: "hello",
	})
	if err != nil {
		panic(err)
	}
	return &resp, nil

}
