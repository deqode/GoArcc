package authentication

import (
	"alfred/modules/authentication/v1/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *authenticationServer) Login(ctx context.Context, in *empty.Empty) (*pb.LoginResponse, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	state := base64.StdEncoding.EncodeToString(b)

	return &pb.LoginResponse{
		Url: s.authenticator.Config.AuthCodeURL(state),
	}, nil
}
