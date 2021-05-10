package authentication

import (
	"alfred/modules/authentication/v1/pb"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = status.Error(codes.NotFound, "No Record Found")

func (s *authenticationServer) LoginCallback(ctx context.Context, in *pb.LoginCallbackRequest) (*pb.LoginCallbackResponse, error) {

	token, err := s.authenticator.Config.Exchange(ctx, in.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, status.Error(codes.Internal, "No id_token field in oauth2 token")
	}

	// TODO: Move this code outside of function
	oidcConfig := &oidc.Config{
		ClientID: s.config.Auth.Auth0ClientID,
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

	//get user_details
	isCreateUserProfile := false
	userID := ""
	usr, err := s.userProfileClient.GetUserProfileBySub(ctx, &userProfilePb.GetUserProfileBySubRequest{
		Sub: fmt.Sprintf("%s", profile["sub"]),
	})
	if err != nil {
		code, _ := status.FromError(err)
		if code.Code() == codes.NotFound {
			isCreateUserProfile = true
		} else {
			return nil, err
		}
	}
	//store user details
	if isCreateUserProfile {
		user, err := s.userProfileClient.CreateUserProfile(ctx, &userProfilePb.CreateUserProfileRequest{
			UserProfile: &userProfilePb.UserProfile{
				Id:             "",
				Sub:            fmt.Sprintf("%s", profile["sub"]),
				Name:           fmt.Sprintf("%s", profile["name"]),
				UserName:       fmt.Sprintf("%s", profile["nickname"]),
				ProfilePicUrl:  fmt.Sprintf("%s", profile["picture"]),
				ExternalSource: userProfilePb.SOURCE_GITHUB,
				TokenValidTill: nil,
			},
		})
		if err != nil {
			return nil, err
		}
		userID = user.Id
	} else {
		userID = usr.Id
	}

	return &pb.LoginCallbackResponse{
		IdToken:     rawIDToken,
		AccessToken: token.AccessToken,
		UserId:      userID,
	}, nil
}
