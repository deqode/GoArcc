package UserProfileService_test

import (
	"alfred/client/grpcClient"
	config "alfred/config"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	"context"
	"log"
	"testing"
)

//As our interfaces has external dependencies
// so we are using go mock when we have to test methods , receivers method
// ~ Mocking simply means resolving all the dependency so that we can test easily.
func TestCreateUserProfile(t *testing.T) {
	clientCon := grpcClient.GetGrpcClientConnection(config.GetConfig())
	client := userProfilePb.NewUserProfileServiceClient(clientCon)
	resp, err := client.CreateUserProfile(, &userProfilePb.CreateUserProfileRequest{UserProfile: &userProfilePb.UserProfile{
		Name:        "Atul Anand",
		Email:       "aanand@deqode.com",
		PhoneNumber: "7717703727",
		ExternalSource: userProfilePb.SOURCE_GITHUB,
	}})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

}
