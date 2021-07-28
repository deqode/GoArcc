package external_svc

import (
	"context"
	"errors"
	"github.com/deqode/GoArcc/config"
	model "github.com/deqode/GoArcc/modules/user/v1/models"
	"github.com/deqode/GoArcc/modules/user/v1/pb"
	"github.com/deqode/GoArcc/util/userinfo"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type userServer struct {
	db     *gorm.DB
	config *config.Config
}

func NewUserServer(
	db *gorm.DB,
	config *config.Config,
) pb.UsersServer {
	model.InitialMigrationUser(db)
	return &userServer{
		db:     db,
		config: config,
	}
}

func (s *userServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	//user model
	var user model.User
	usrId := userinfo.FromContext(ctx).ID
	gormDb := s.db
	//ie: Select * from account where id = in.id
	if err := gormDb.First(&user, "id= ?", usrId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &pb.User{
		Id:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		UserName:      user.UserName,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		ProfilePicUrl: user.ProfilePicURL,
	}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	return nil, nil
}

func (s *userServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*empty.Empty, error) {
	//request validation
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request can't be nil")
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//by default it will delete with primary key.
	// ie: Delete From Account where id = in.id
	tx := s.db.Where("id = ?", in.Id).Delete(&model.User{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &empty.Empty{}, nil
}
