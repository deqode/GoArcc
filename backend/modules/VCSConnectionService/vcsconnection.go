package VCSConnectionService

import (
	"alfred/modules/VCSConnectionService/pb"
	"context"
	"github.com/coreos/go-oidc"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func (s *VCSConnectionService) ListAlfredVCSConnection(context.Context, *empty.Empty) (*pb.ListAlfredVCSConnectionResponse, error) {
	return nil, nil
}
func (s *VCSConnectionService) Connect(context.Context, *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	return nil, nil
}
func (s *VCSConnectionService) Connected(context.Context, *pb.ConnectedRequest) (*pb.AccountVCSConnection, error) {
	return nil, nil
}
func (s *VCSConnectionService) ListVCSConnection(context.Context, *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	return nil, nil
}
func (s *VCSConnectionService) RevokeVCSToken(context.Context, *pb.RevokeVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *VCSConnectionService) RenewVCSToken(context.Context, *pb.RenewVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
