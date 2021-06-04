package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) CreateStack(ctx context.Context, in *pb.CreateStackRequest) (*pb.Stack, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, in.Stack.AccountId, s.db); err != nil {
		return nil, err
	}
	stackModel := model.PrepareStackModel(in.Stack)
	//insert into db
	if s.db.Create(stackModel).Error != nil {
		return nil, status.Error(codes.Internal, "Record not inserted in Stack")
	}
	resStack := model.PrepareStackObj(stackModel)
	return resStack, nil
}
