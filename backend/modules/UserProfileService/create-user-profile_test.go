package UserProfileService_test

import (
	"alfred/client/grpcClient"
	config "alfred/config"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"log"
	"testing"
)

//As our interfaces has external dependencies
// so we are using go mock when we have to test methods , receivers method
// ~ Mocking simply means resolving all the dependency so that we can test easily.
func TestCreateUserProfile(t *testing.T) {
	clientContext := grpcClient.GetGrpcClientConnectionContext()
	conn := grpcClient.GetGrpcClientConnection(config.GetConfig(), clientContext)
	defer conn.Close()
	client := userProfilePb.NewUserProfileServiceClient(conn)
	resp, err := client.CreateUserProfile(clientContext.Ctx, &userProfilePb.CreateUserProfileRequest{UserProfile: &userProfilePb.UserProfile{
		Name:        "Atul Anand",
		Email:       "aanand@deqode.com",
		PhoneNumber: "7717703727",
		Source:      0,
	}})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

}
