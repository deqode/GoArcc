package external_svc

import (
	"alfred/modules/stack/v1/internal-svc"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) ListStackBuild(ctx context.Context, in *pb.ListStackBuildRequest) (*pb.ListStackBuildResponse, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// Validate Stack
	st := &model.Stack{ID: in.StackId}
	if err := s.db.First(st).Error; err != nil {
		return nil, status.Error(codes.NotFound, "No Stack Found")
	}

	// validate User
	if err := internal_svc.ValidateUserFromContext(ctx, st.AccountID, s.db); err != nil {
		return nil, err
	}

	var stackBuildModels []*model.StackBuild
	tx := s.db.Find(&stackBuildModels).Where("stack_id = ?", in.GetStackId())
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}

	var stacksBuildRes []*pb.StackBuild
	for _, st := range stackBuildModels {
		stacksBuildRes = append(stacksBuildRes, model.PrepareStackBuildObj(st))
	}

	return &pb.ListStackBuildResponse{
		StackBuilds: stacksBuildRes,
	}, nil
}
