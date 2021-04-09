package AuthService

import (
	"alfred/config"
	"alfred/modules/AuthService/v1/pb"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/coreos/go-oidc"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/url"
	//userProfilePb "alfred/modules/UserProfileService/pb"
)

func (s *AuthService) UserLogin(ctx context.Context, in *empty.Empty) (*pb.UserLoginResponse, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	state := base64.StdEncoding.EncodeToString(b)

	return &pb.UserLoginResponse{
		Url: s.authenticator.Config.AuthCodeURL(state),
	}, nil
}

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

func (s *AuthService) UserLogout(context.Context, *empty.Empty) (*empty.Empty, error) {
	//TODO - Implement Token based logout
	domain := config.GetConfig().Auth.Auth0Domain

	logoutUrl, err := url.Parse("http://" + domain)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	logoutUrl.Path += "/v1/authentication/logout"
	parameters := url.Values{}

	var scheme string
	scheme = "http"

	returnTo, err := url.Parse(scheme + "://" + "http://localhost:8082")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", config.GetConfig().Auth.Auth0ClientId)
	logoutUrl.RawQuery = parameters.Encode()

	return &empty.Empty{}, nil
}
