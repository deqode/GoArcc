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
	c := config.Config{Grpc: config.GraphqlServerConfig{
		Port: "8080",
		Host: "localhost",
	}}
	clientCon := grpcClient.GetGrpcClientConnection(&c)
	client := userProfilePb.NewUserProfileServiceClient(clientCon)
	resp, err := client.CreateUserProfile(context.Background(), &userProfilePb.CreateUserProfileRequest{UserProfile: &userProfilePb.UserProfile{
		Name:           "Shivang Saxena",
		Email:          "ssaxena@deqode.com",
		PhoneNumber:    "",
		ExternalSource: userProfilePb.SOURCE_GITHUB,
	}})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

}

func TestDeleteUserProfile(t *testing.T) {
	c := config.Config{Grpc: config.GraphqlServerConfig{
		Port: "8080",
		Host: "localhost",
	}}
	clientCon := grpcClient.GetGrpcClientConnection(&c)
	client := userProfilePb.NewUserProfileServiceClient(clientCon)
	resp, err := client.DeleteUserProfile(context.Background(), &userProfilePb.DeleteUserProfileRequest{Id: "c692a624-077b-c992-bd49-c9d06c48e717"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

}

func TestGetUserProfile(t *testing.T) {
	c := config.Config{Grpc: config.GraphqlServerConfig{
		Port: "8080",
		Host: "localhost",
	}}
	clientCon := grpcClient.GetGrpcClientConnection(&c)
	client := userProfilePb.NewUserProfileServiceClient(clientCon)
	resp, err := client.GetUserProfile(context.Background(), &userProfilePb.GetUserProfileRequest{Id: "5531657c-9e9e-ddc5-e4f6-30b2779a71a9"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

}
