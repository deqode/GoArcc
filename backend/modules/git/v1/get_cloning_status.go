package git

import (
	"alfred/modules/git/v1/pb"
	"context"
)

func (s *gitServer) GetCloningStatus(ctx context.Context, in *pb.GetCloningStatusRequest) (*pb.GetCloningStatusResponse, error) {
	//workflow := s.cadenceAdapter.CadenceClient.GetWorkflow(ctx, in.WorkflowId, in.RunId)
	//var result pb.GetCloningStatusResponse
	////todo : handle error @shivang
	//workflow.Get(ctx, &result)
	return nil, nil
}
