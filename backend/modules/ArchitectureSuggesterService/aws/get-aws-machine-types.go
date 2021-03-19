package aws

import (
	"alfred/modules/ArchitectureSuggesterService/aws/awsDefault"
	"alfred/modules/ArchitectureSuggesterService/aws/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (a aws.AwsService) GetAwsMachineTypes(ctx context.Context, request *pb.GetAwsMachineTypesRequest) (*pb.GetAwsMachineTypesResponse, error) {
	machineTypes := make([]string, 0)
	switch request.InstanceType {
	case "GENERAL_PURPOSE":
		machineTypes = awsDefault.GetAwsMachineTypes(ctx)
	case "COMPUTE_OPTIMIZED":
		//todo we will implement it later
	case "MEMORY_OPTIMIZED":
		//todo we will implement it later
	case "ACCELERATED_COMPUTING":
		//todo we will implement it later
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid launch type or launch type not exist")
	}
	return &pb.GetAwsMachineTypesResponse{MachineTypes: machineTypes}, nil
}
