package AuthService

import (
	"alfred/config"
	"alfred/modules/AuthService/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/coreos/go-oidc"
	empty "github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"net/url"
)

func (s *AuthService) UserLogin(ctx context.Context, in *empty.Empty) (*pb.UserLoginResponse, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	state := base64.StdEncoding.EncodeToString(b)

	//session, err := Store.Get(r, "auth-session")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//session.Values["state"] = state
	//err = session.Save(r, w)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	authenticator, err := NewAuthenticator(config.GetConfig())
	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		Url: authenticator.Config.AuthCodeURL(state),
	}, nil
}
func (s *AuthService) UserLoginCallback(ctx context.Context, in *pb.UserLoginCallbackRequest) (*empty.Empty, error) {
	//session, err := Store.Get(r, "auth-session")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//if r.URL.Query().Get("state") != session.Values["state"] {
	//	http.Error(w, "Invalid state parameter", http.StatusBadRequest)
	//	return
	//}

	authenticator, err := NewAuthenticator(config.GetConfig())
	if err != nil {
		return nil, err
	}

	token, err := authenticator.Config.Exchange(context.TODO(), in.Code)
	if err != nil {
		//log.Printf("no token found: %v", err)
		//w.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		//return err http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return nil, http.ErrAbortHandler
	}

	oidcConfig := &oidc.Config{
		ClientID: config.GetConfig().Auth.Auth0ClientId,
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		//return http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, err
	}

	//session.Values["id_token"] = rawIDToken
	//session.Values["access_token"] = token.AccessToken
	//session.Values["profile"] = profile
	//err = session.Save(r, w)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

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
