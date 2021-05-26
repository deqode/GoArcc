package internal_svc

import (
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes"
)

func (s *vcsConnectionIntServer) CreateVCSConnection(ctx context.Context, in *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error) {
	//validate a request
	if err := in.Validate(); err != nil {
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
	in.VcsConnection.Id = VCSModel.ID
	return in.VcsConnection, nil
}
