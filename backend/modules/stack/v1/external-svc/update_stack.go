package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) UpdateStack(ctx context.Context, in *pb.UpdateStackRequest) (*pb.Stack, error) {
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request is nil")
	}
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, in.GetStack().GetAccountId(), s.db); err != nil {
		return nil, err
	}
	stackModel := model.PrepareStackModel(in.Stack)
	// Omit Exclude fields from updating
	tx := s.db.Model(&stackModel).Omit("id", "account_id", "cloud_connection_id", "vcs_connection_id").Updates(&stackModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.Internal, "Record not updated")
	}

	updatedStack := model.PrepareStackObj(stackModel)
	return updatedStack, nil
}
