package AuthService

import (
	"alfred/modules/AuthService/v1/pb"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthService) UserLoginCallback(ctx context.Context, in *pb.UserLoginCallbackRequest) (*pb.UserLoginCallbackResponse, error) {

	token, err := s.authenticator.Config.Exchange(ctx, in.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, status.Error(codes.Internal, "No id_token field in oauth2 token")
	}
	oidcConfig := &oidc.Config{
		ClientID: s.config.Auth.Auth0ClientId,
	}

	idToken, err := s.authenticator.Provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to verify ID Token: "+err.Error())
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//store user details
	usr, err := s.userProfileClient.CreateUserProfile(ctx, &userProfilePb.CreateUserProfileRequest{
		UserProfile: &userProfilePb.UserProfile{
			Id:             "",
			Sub:            fmt.Sprintf("%s", profile["sub"]),
			Name:           fmt.Sprintf("%s", profile["nickname"]),
			Email:          fmt.Sprintf("%s", profile["name"]),
			PhoneNumber:    "",
			ExternalSource: userProfilePb.SOURCE_GITHUB,
			TokenValidTill: nil,
		},
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserLoginCallbackResponse{
		IdToken:     rawIDToken,
		AccessToken: token.AccessToken,
		UserId:      usr.Id,
	}, nil
}