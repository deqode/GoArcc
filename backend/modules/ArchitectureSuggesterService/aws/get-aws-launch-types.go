package aws

import (
	"alfred/modules/ArchitectureSuggesterService/aws/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a AwsService) GetAwsLaunchTypes(ctx context.Context, request *pb.GetAwsLaunchTypesRequest) (*pb.GetAwsLaunchTypesResponse, error) {

	if a.awsDefaultData == nil ||
		a.awsDefaultData.AwsContainerServices == nil {
		return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting aws container services")
	}

	launchType := make([]string, 0)
	switch request.ContainerService {
	case "ECS":
		if a.awsDefaultData.AwsContainerServices.EcsContainerService == nil ||
			a.awsDefaultData.AwsContainerServices.EcsContainerService.Name == nil ||
			len(a.awsDefaultData.AwsContainerServices.EcsContainerService.Name) == 0 {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting ecs launch types")
		}
		launchType = a.awsDefaultData.AwsContainerServices.EcsContainerService.Name
	case "EKS":
		if a.awsDefaultData.AwsContainerServices.EksContainerService == nil ||
			a.awsDefaultData.AwsContainerServices.EksContainerService.Name == nil ||
			len(a.awsDefaultData.AwsContainerServices.EksContainerService.Name) == 0 {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting eks launch types")
		}
		launchType = a.awsDefaultData.AwsContainerServices.EksContainerService.Name
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid container service or container service not exist")
	}
	return &pb.GetAwsLaunchTypesResponse{LaunchType: launchType}, nil
}
