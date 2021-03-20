package aws

import (
	awsDefaultDataPb "alfred/modules/ArchitectureSuggesterService/aws/awsDefault/pb"
	"alfred/modules/ArchitectureSuggesterService/aws/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a AwsService) GetAwsInstanceTypes(ctx context.Context, request *pb.GetAwsInstanceTypesRequest) (*pb.GetAwsInstanceTypesResponse, error) {

	if a.awsDefaultData == nil ||
		a.awsDefaultData.AwsContainerServices == nil {
		return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting aws container service")
	}

	instanceTypes := make([]string, 0)

	switch request.Parent {
	case "ECS":
		if a.awsDefaultData.AwsContainerServices.EcsContainerService == nil {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting ecs container service")
		}
		switch request.LaunchType {
		case "EC2":
			if err := validate(request.LaunchType, a.awsDefaultData); err != nil {
				return nil, err
			}
			instanceTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.Name
		case "FRAGRATE":
			if err := validate(request.LaunchType, a.awsDefaultData); err != nil {
				return nil, err
			}
			instanceTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.FragrateLaunchType.Name
		default:
			return nil, status.Error(codes.InvalidArgument, "invalid launch type or launch type not exist")
		}

	case "EKS":
		if a.awsDefaultData.AwsContainerServices.EksContainerService == nil {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting eks container service")
		}
		switch request.LaunchType {
		//todo will implemented later
		default:
			return nil, status.Error(codes.InvalidArgument, "invalid launch type or launch type not exist")
		}

	}

	return &pb.GetAwsInstanceTypesResponse{InstanceTypes: instanceTypes}, nil
}

func validate(launchType string, data *awsDefaultDataPb.AwsDefaultData) error {

	switch launchType {
	case "EC2":
		if data.AwsContainerServices.EcsContainerService.Ec2LaunchType == nil ||
			data.AwsContainerServices.EcsContainerService.Ec2LaunchType.Name == nil ||
			len(data.AwsContainerServices.EcsContainerService.Ec2LaunchType.Name) == 0 {
			return status.Error(codes.FailedPrecondition, "something went wrong while getting Ec2 instance type")
		}
	case "FRAGRATE":
		if data.AwsContainerServices.EcsContainerService.Ec2LaunchType == nil ||
			data.AwsContainerServices.EcsContainerService.FragrateLaunchType.Name == nil ||
			len(data.AwsContainerServices.EcsContainerService.FragrateLaunchType.Name) == 0 {
			return status.Error(codes.FailedPrecondition, "something went wrong while getting Fragrate instance type")
		}

	}
	return nil
}
