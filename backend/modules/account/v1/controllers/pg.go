package controllers

import (
	model "alfred/modules/account/v1/models"
	"context"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type accountPgStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) AccountStore {
	return &accountPgStore{db}
}

func (s *accountPgStore) CreateAccount(ctx context.Context, in []*model.Account) ([]string, error) {
	//TODO - Validations
	var ids []string
	for _, acc := range in {
		id, err := uuid.GenerateUUID()
		if err != nil {
			return nil, err
		}
		acc.ID = id
		ids = append(ids, id)
	}
	t := s.db.Create(in)
	if t.Error != nil {
		return nil, t.Error
	}
	return ids, nil
}

func (s *accountPgStore) GetAccount(ctx context.Context, condition map[string]string) (*model.Account, error) {
	//TODO - Validations
	account := model.Account{}
	chain := s.db
	for k, v := range condition {
		chain = chain.Where(k, v)
	}
	if err := chain.Find(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &account, nil
}

func (s *accountPgStore) ListAccount(ctx context.Context, condition map[string]string) ([]*model.Account, error) {
	//TODO - Validations
	var accounts []*model.Account
	chain := s.db
	for k, v := range condition {
		chain = chain.Where(k, v)
	}
	if err := chain.Find(&accounts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return accounts, nil
}

func (s *accountPgStore) UpdateAccount(ctx context.Context, in *model.Account, updateMask *field_mask.FieldMask) (*model.Account, error) {
	account := model.Account{}

	return &account, nil
}

func (s *accountPgStore) DeleteAccount(ctx context.Context, condition map[string]string) (bool, error) {
	chain := s.db
	for k, v := range condition {
		chain = chain.Where(k, v)
	}
	//if err := s.db.Delete(&models.Account{ID: in.Id}).Error; err != nil {
	//	return nil, err
	//}
	return false, nil
}
