package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
)

func (s *stacksServer) GetStackBuild(ctx context.Context, in *pb.GetStackBuildRequest) (*pb.StackBuild, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	stackBuildModel := &model.StackBuild{ID: in.Id}
	if err := s.db.First(stackBuildModel).Error; err != nil {
		return nil, err
	}

	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, stackBuildModel.AccountID, s.db); err != nil {
		return nil, err
	}

	stackRes := model.PrepareStackBuildObj(stackBuildModel)
	return stackRes, nil
}
