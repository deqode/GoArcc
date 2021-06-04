package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) ListStack(ctx context.Context, in *pb.ListStackRequest) (*pb.ListStackResponse, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, in.AccountId, s.db); err != nil {
		return nil, err
	}

	var stackModels []*model.Stack
	tx := s.db.Find(&stackModels).Where("account_id = ?", in.GetAccountId())
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}

	var stacksRes []*pb.Stack
	for _, st := range stackModels {
		stacksRes = append(stacksRes, model.PrepareStackObj(st))
	}

	return &pb.ListStackResponse{
		Stacks: stacksRes,
	}, nil
}
