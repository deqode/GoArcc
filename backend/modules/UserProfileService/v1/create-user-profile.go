package v1

import (
	"alfred/logger"
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"context"
	"fmt"
	"gorm.io/gorm"
	"os/exec"
	"time"
)

func (s *UserProfileService) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	//creating uuid
	out, err := exec.Command("uuidgen").Output() //
	if err != nil {
		logger.Log.Debug("unable to generate uuid")
	}

	VCSModel := &models.UserProfile{
		ID:          fmt.Sprintf("%s", out),
		Name:        in.UserProfile.Name,
		Email:       in.UserProfile.Email,
		PhoneNumber: in.UserProfile.PhoneNumber,
		Sub:         in.UserProfile.Sub,
		Source:      1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}

	t := s.db.Create(VCSModel)
	if t.Error != nil {
		return nil, t.Error
	}

	//TODO - create a single account as well
	return in.UserProfile, nil
}
