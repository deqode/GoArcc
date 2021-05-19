package models

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type userProfilePgStore struct {
	db *gorm.DB
}

func NewUserProfileStore(db *gorm.DB) UserProfileStore {
	return &userProfilePgStore{db}
}

func (s userProfilePgStore) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	_, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	usrProfile := &UserProfile{
		ID:            in.UserProfile.Sub,
		Name:          in.UserProfile.Name,
		Email:         in.UserProfile.Email,
		PhoneNumber:   in.UserProfile.PhoneNumber,
		Sub:           in.UserProfile.Sub,
		UserName:      in.UserProfile.UserName,
		ProfilePicURL: in.UserProfile.ProfilePicUrl,
		Source:        in.UserProfile.ExternalSource,
	}

	t := s.db.Create(usrProfile)
	if t.Error != nil {
		return nil, t.Error
	}

	return in.UserProfile, nil
}

func (s userProfilePgStore) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	profile := UserProfile{}
	result := s.db.First(&profile, "id = ?", in.Id)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &pb.UserProfile{
		Id:             profile.ID,
		ExternalSource: profile.Source,
		Name:           profile.Name,
		Email:          profile.Email,
		Sub:            profile.Sub,
		UserName:       profile.UserName,
	}, nil
}

func (s userProfilePgStore) GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error) {
	profile := UserProfile{}
	err := s.db.First(&profile, "sub = ?", in.Sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &pb.UserProfile{
		Id:             profile.ID,
		ExternalSource: profile.Source,
		Name:           profile.Name,
		Email:          profile.Email,
		Sub:            profile.Sub,
		UserName:       profile.UserName,
	}, nil
}

func (s userProfilePgStore) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}

func (s userProfilePgStore) DeleteUserProfile(ctx context.Context, request *pb.DeleteUserProfileRequest) (*empty.Empty, error) {

	// this will be soft delete from the system
	t := s.db.Delete(&UserProfile{
		ID: request.Id,
	})
	if t.Error != nil {
		return nil, t.Error
	}
	return nil, nil
}
