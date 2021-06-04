package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) CreateStackBuild(ctx context.Context, in *pb.CreateStackBuildRequest) (*pb.StackBuild, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, in.StackBuild.AccountId, s.db); err != nil {
		return nil, err
	}
	stackBuildModel := model.PrepareStackBuildModel(in.StackBuild)
	//insert into db
	if s.db.Create(stackBuildModel).Error != nil {
		return nil, status.Error(codes.Internal, "Record not inserted in Stack")
	}
	resStack := model.PrepareStackBuildObj(stackBuildModel)
	return resStack, nil
}
