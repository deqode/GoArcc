package internal_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *userProfileInServer) DeleteUserProfile(ctx context.Context, in *pb.DeleteUserProfileRequest) (*empty.Empty, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//by default it will delete with primary key.
	// ie: Delete From Account where id = in.id
	tx := s.db.Delete(&model.UserProfile{}, in.Id)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
