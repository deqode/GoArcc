package AccountService

import (
	"alfred/modules/AccountService/v1/models"
	"alfred/modules/AccountService/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/hashicorp/go-uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AccountService) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	if len(in.Account.UserId) < 3 {
		return nil, status.Error(codes.FailedPrecondition, "UserId is required")
	}
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	account := &models.Account{
		ID:     id,
		Slug:   in.GetAccount().GetSlug(),
		UserId: in.GetAccount().GetUserId(),
	}
	t := s.db.Create(account)
	if t.Error != nil {
		return nil, t.Error
	}
	in.Account.Id = id
	return in.Account, nil
}

func (s *AccountService) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	account := &models.Account{
		Slug: in.GetAccount().GetSlug(),
	}
	t := s.db.Model(&account).Where("id ="+in.Account.Id, true).Update("slug", in.Account.Slug)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.Account, nil
}

func (s *AccountService) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	account := pb.Account{}
	result := s.db.First(&account, "id = ?", in.Id)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &account, nil
}

func (s *AccountService) DeleteAccount(context.Context, *pb.DeleteAccountRequest) (*empty.Empty, error) {
	return nil, nil
}
