package UserProfileService

import (
	accountPb "alfred/modules/Account/v1/pb"
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"context"
	"github.com/hashicorp/go-uuid"
)

func (s *UserProfileService) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	VCSModel := &models.UserProfile{
		ID:            id,
		Name:          in.UserProfile.Name,
		Email:         in.UserProfile.Email,
		PhoneNumber:   in.UserProfile.PhoneNumber,
		Sub:           in.UserProfile.Sub,
		UserName:      in.UserProfile.UserName,
		ProfilePicURL: in.UserProfile.ProfilePicUrl,
		Source:        in.UserProfile.ExternalSource,
	}

	t := s.db.Create(VCSModel)
	if t.Error != nil {
		return nil, t.Error
	}

	// create a user account as well
	_, err = s.accountClient.CreateAccount(ctx, &accountPb.CreateAccountRequest{
		Account: &accountPb.Account{
			Slug:   VCSModel.Name + "-" + in.UserProfile.ExternalSource.String(),
			UserId: id,
		},
	})
	if err != nil {
		return nil, err
	}
	in.UserProfile.Id = id
	return in.UserProfile, nil
}
