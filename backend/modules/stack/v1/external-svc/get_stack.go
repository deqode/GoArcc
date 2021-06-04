package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
)

func (s *stacksServer) GetStack(ctx context.Context, in *pb.GetStackRequest) (*pb.Stack, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	stackModel := &model.Stack{ID: in.Id}
	if err := s.db.First(stackModel).Error; err != nil {
		return nil, err
	}

	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, stackModel.AccountID, s.db); err != nil {
		return nil, err
	}

	stackRes := model.PrepareStackObj(stackModel)
	return stackRes, nil
}
