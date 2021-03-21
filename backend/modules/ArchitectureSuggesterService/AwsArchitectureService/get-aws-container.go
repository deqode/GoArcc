package AwsArchitectureService

import (
	"alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a AwsArchitectureService) GetAwsContainerServices(ctx context.Context, request *pb.GetAwsContainerServicesRequest) (*pb.GetAwsContainerServicesResponse, error) {

	if a.awsDefaultData != nil {
		return nil, status.Error(codes.FailedPrecondition, "Aws configuration not found")
	}
	if a.awsDefaultData.AwsContainerServices != nil ||
		a.awsDefaultData.AwsContainerServices.Name != nil ||
		len(a.awsDefaultData.AwsContainerServices.Name) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "Aws container services not found")
	}
	return &pb.GetAwsContainerServicesResponse{AmazonContainerServices: a.awsDefaultData.AwsContainerServices.Name}, nil
}
