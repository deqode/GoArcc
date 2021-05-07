package AuthService

import (
	"alfred/modules/auth/v1/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang/protobuf/ptypes/empty"
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
