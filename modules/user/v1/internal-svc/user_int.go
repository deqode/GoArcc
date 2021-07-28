package internal_svc

import (
	"context"
	"github.com/deqode/GoArcc/config"
	model "github.com/deqode/GoArcc/modules/user/v1/models"
	"github.com/deqode/GoArcc/modules/user/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

// Internal Service Configuration
type userInServer struct {
	db     *gorm.DB
	config *config.Config
}

func NewUserInServer(
	db *gorm.DB,
	config *config.Config,
) pb.UserInternalServer {
	model.InitialMigrationUser(db)
	return &userInServer{
		db:     db,
		config: config,
	}
}

func (s *userInServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	if in.User == nil || in == nil {
		return nil, status.Error(codes.FailedPrecondition, "User not provided")
	}
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	//prepare insert object
	UserModel := &model.User{
		ID:            in.GetUser().GetId(),
		FirstName:     in.GetUser().GetFirstName(),
		LastName:      in.GetUser().GetLastName(),
		UserName:      in.GetUser().GetUserName(),
		Email:         in.GetUser().GetEmail(),
		PhoneNumber:   in.GetUser().GetPhoneNumber(),
		ProfilePicURL: in.GetUser().GetProfilePicUrl(),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
	//insert into db
	gormDb := s.db
	tx := gormDb.Create(UserModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return in.User, nil
}
