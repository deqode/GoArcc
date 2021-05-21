package internal_svc

import (
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/hashicorp/go-uuid"
)

func (s *vcsConnectionIntServer) CreateVCSConnection(ctx context.Context, in *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error) {
	// TODO - validate user belongs to that user account
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
	VCSModel := &models.VCSConnection{
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
