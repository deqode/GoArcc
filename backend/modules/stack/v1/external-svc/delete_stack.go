package external_svc

import (
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *stacksServer) DeleteStack(ctx context.Context, in *pb.DeleteStackRequest) (*empty.Empty, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	tx := s.db.Delete(&model.Stack{ID: in.Id})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.Internal, "Record not deleted")
	}
	return &empty.Empty{}, nil
}
