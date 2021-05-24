package internal_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userProfileInServer) DeleteUserProfile(ctx context.Context, in *pb.DeleteUserProfileRequest) (*empty.Empty, error) {
	//request validation
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request can't be nil")
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//by default it will delete with primary key.
	// ie: Delete From Account where id = in.id
	tx := s.db.Where("id = ?", in.Id).Delete(&model.UserProfile{})
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
