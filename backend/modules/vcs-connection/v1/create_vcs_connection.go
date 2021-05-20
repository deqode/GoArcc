package vcs_connection

import (
	models2 "alfred/modules/vcs-connection/v1/internal/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/hashicorp/go-uuid"
)

func (s *vcsConnectionServer) CreateVCSConnection(ctx context.Context, in *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	aTEP, err := ptypes.Timestamp(in.VcsConnection.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}
	rTEP, err := ptypes.Timestamp(in.VcsConnection.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}
	VCSModel := &models2.VCSConnection{
		ID:                 id,
		Provider:           in.VcsConnection.Provider,
		ConnectionID:       in.VcsConnection.ConnectionId,
		AccessToken:        in.VcsConnection.AccessToken,
		RefreshToken:       in.VcsConnection.RefreshToken,
		AccessTokenExpiry:  &aTEP,
		RefreshTokenExpiry: &rTEP,
		Revoked:            false,
		AccountID:          in.VcsConnection.AccountId,
		UserName:           in.VcsConnection.UserName,
		Label:              in.VcsConnection.Label,
	}

	t := s.db.Create(VCSModel)
	if t.Error != nil {
		return nil, t.Error
	}
	//todo Get from DB
	return in.VcsConnection, nil
}
