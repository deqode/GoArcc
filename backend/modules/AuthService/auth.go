package AuthService

import (
	"alfred/config"
	"alfred/modules/AuthService/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/oauth2"
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

	authenticator, err := NewAuthenticator(config.GetConfig())
	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		Url: authenticator.Config.AuthCodeURL(state),
	}, nil
}
func (s *AuthService) UserLoginCallback(ctx context.Context, in *pb.UserLoginCallbackRequest) (*empty.Empty, error) {

	authenticator, err := NewAuthenticator(config.GetConfig())
	if err != nil {
		return nil, err
	}

	token, err := authenticator.Config.Exchange(context.TODO(), in.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, status.Error(codes.Internal, "No id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: config.GetConfig().Auth.Auth0ClientId,
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to verify ID Token: "+err.Error())
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//store user details
	//_, err = s.userClient.CreateUserProfile(ctx, &userProfilePb.CreateUserProfileRequest{
	//	UserProfile: &userProfilePb.UserProfile{
	//		Id:             "",
	//		Sub:            fmt.Sprintf("%s", profile["sub"]),
	//		Name:           fmt.Sprintf("%s", profile["nickname"]),
	//		Email:          fmt.Sprintf("%s", profile["name"]),
	//		PhoneNumber:    "",
	//		Source:         userProfilePb.SOURCE_GITHUB,
	//		TokenValidTill: nil,
	//	},
	//})
	//if err != nil {
	//	return nil, err
	//}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, status.Error(codes.Internal, err.Error())
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, status.Error(codes.Internal, err.Error())
	}
	fmt.Println(data)
	// Redirect to logged in page
	return &empty.Empty{}, nil
}
func (s *AuthService) UserLogout(context.Context, *empty.Empty) (*empty.Empty, error) {
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
	//if r.TLS == nil {
	//	scheme = "http"
	//} else {
	//	scheme = "https"
	//}

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

// GetUserLogin returns the specified user by its id.
func (s *AuthService) GetUserLogin(context.Context, *pb.GetUserLoginRequest) (*empty.Empty, error) {
	return nil, nil
}

// DeleteUserLogin is used to delete a user from the system, this will delete all
func (s *AuthService) DeleteUserLogin(context.Context, *pb.DeleteUserLoginRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *AuthService) UpdateUserPassword(context.Context, *pb.UpdateUserPasswordRequest) (*empty.Empty, error) {
	return nil, nil
}

// ResetUserPassword , if a user has forgot the password then
// this method can be used to reset the password
func (s *AuthService) ResetUserPassword(context.Context, *pb.ResetUserPasswordRequest) (*empty.Empty, error) {
	return nil, nil
}
