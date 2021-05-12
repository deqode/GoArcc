package authentication

import (
	accountInPb "alfred/modules/account/v1/internals/pb"
	"alfred/modules/authentication/v1/pb"
	userProfileInPb "alfred/modules/user-profile/v1/internals/pb"
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
	usr, err := s.userProfileServer.GetUserProfileBySub(ctx, &userProfilePb.GetUserProfileBySubRequest{
		Sub: fmt.Sprintf("%s", profile["sub"]),
	})
	if usr != nil {
		return &pb.LoginCallbackResponse{
			IdToken:     rawIDToken,
			AccessToken: token.AccessToken,
			UserId:      usr.Id,
		}, nil
	}
	if err != nil {
		return nil, err
	}

	//create User Profile
	user, err := s.userProfileInServer.CreateUserProfile(ctx, &userProfileInPb.CreateUserProfileRequest{
		UserProfile: &userProfileInPb.UserProfile{
			Id:             "",
			Sub:            fmt.Sprintf("%s", profile["sub"]),
			Name:           fmt.Sprintf("%s", profile["name"]),
			UserName:       fmt.Sprintf("%s", profile["nickname"]),
			ProfilePicUrl:  fmt.Sprintf("%s", profile["picture"]),
			ExternalSource: userProfileInPb.SOURCE_GITHUB,
			TokenValidTill: nil,
		},
	})
	if err != nil {
		return nil, err
	}

	//create User Account
	_, err = s.accountServer.CreateAccount(ctx, &accountInPb.CreateAccountRequest{
		Account: &accountInPb.Account{
			Slug:   user.Name + "_" + user.ExternalSource.String(),
			UserId: user.Id,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginCallbackResponse{
		IdToken:     rawIDToken,
		AccessToken: token.AccessToken,
		UserId:      usr.Id,
	}, nil
}
