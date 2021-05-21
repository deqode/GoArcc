package internal_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserProfileStore interface {
	GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.UserProfile, error)
	CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error)
	UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.UserProfile, error)
	DeleteUserProfile(ctx context.Context, in *pb.DeleteUserProfileRequest) (*empty.Empty, error)
	GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error)
}
