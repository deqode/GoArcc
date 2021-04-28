package VCSConnectionService

import (
	"alfred/background/workers/workflows"
	"alfred/background/workers/workflows/vcs/git"
	"alfred/logger"
	"context"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)


//trigger
func (s *VCSConnectionService) triggerGitClone(req git.GitCloneRequest) error {
	wo := client.StartWorkflowOptions{
		TaskList:                     workflows.GitCloneGroup,
		ExecutionStartToCloseTimeout: time.Hour * 24,
	}
	execution, err := s.cadenceAdapter.CadenceClient.StartWorkflow(context.Background(), wo, workflows.GitCloneGroup, &req)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	logger.Log.Info("Started work flow!", zap.String("WorkflowId", execution.ID), zap.String("RunId", execution.RunID))
	return nil
}
