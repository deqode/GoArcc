package external_svc

import (
	"alfred.sh/common/logger"
	"alfred/background/workers/workflows"
	"alfred/background/workers/workflows/vcs/git"
	"alfred/modules/git/v1/pb"
	"context"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *gitServer) CloneRepository(ctx context.Context, in *pb.CloneRepositoryRequest) (*pb.CloneRepositoryResponse, error) {
	//use wait group to allow process to be concurrent
	time.Sleep(time.Second * 10)
	wo := client.StartWorkflowOptions{
		TaskList:                     workflows.GitCloneGroup,
		ExecutionStartToCloseTimeout: time.Hour * 24,
	}
	execution, err := s.cadenceAdapter.CadenceClient.StartWorkflow(ctx, wo, git.GitCloneWorkflow, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.Log.Info("Started work flow!", zap.String("WorkflowId", execution.ID), zap.String("RunId", execution.RunID))
	return &pb.CloneRepositoryResponse{
		RunId:      execution.RunID,
		WorkflowId: execution.ID,
	}, nil
}
