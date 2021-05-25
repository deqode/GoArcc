package external_svc

import (
	"alfred/modules/authentication/v1/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authenticationServer) Login(context.Context, *empty.Empty) (*pb.LoginResponse, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	state := base64.StdEncoding.EncodeToString(b)
	if s.authenticator == nil {
		return nil, status.Error(codes.Internal, "Authenticator not initialised")
	}
	url := s.authenticator.Config.AuthCodeURL(state)

	return &pb.LoginResponse{
		Url: url,
	}, nil
}
