package internal_svc

import (
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
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &empty.Empty{}, nil
}
