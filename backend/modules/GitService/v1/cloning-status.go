package GitService

import (
	"alfred/modules/GitService/v1/pb"
	"context"
)

func (s *GitService) GetCloningStatus(ctx context.Context, in *pb.GetCloningStatusRequest) (*pb.GetCloningStatusResponse, error) {
	workflow := s.cadenceAdapter.CadenceClient.GetWorkflow(ctx, in.WorkflowId, in.RunId)
	var result pb.GetCloningStatusResponse
	workflow.Get(ctx, &result)
	return &result, nil
}
