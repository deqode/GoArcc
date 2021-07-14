package external_svc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"net/url"
)

func (s *authenticationServer) Logout(context.Context, *empty.Empty) (*empty.Empty, error) {
	domain := s.config.Auth.Auth0Domain

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
	parameters.Add("client_id", s.config.Auth.Auth0ClientID)
	logoutURL.RawQuery = parameters.Encode()

	return &empty.Empty{}, nil
}
