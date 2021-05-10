package authentication

import (
	"alfred/config"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"net/url"
)

func (s *authenticationServer) Logout(context.Context, *empty.Empty) (*empty.Empty, error) {
	//TODO - Implement Token based logout
	domain := config.GetConfig().Auth.Auth0Domain

	logoutURL, err := url.Parse("http://" + domain)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	logoutURL.Path += "/v1/authenticationServer/logout"
	parameters := url.Values{}

	var scheme string
	scheme = "http"

	returnTo, err := url.Parse(scheme + "://" + "http://localhost:8082")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", config.GetConfig().Auth.Auth0ClientID)
	logoutURL.RawQuery = parameters.Encode()

	return &empty.Empty{}, nil
}
