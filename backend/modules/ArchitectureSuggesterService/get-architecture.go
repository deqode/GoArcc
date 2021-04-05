package v1

import (
	"alfred/modules/ArchitectureSuggesterService/v1/pb"
	"context"
)

func (a ArchitectureSuggesterService) GetArchitecture(ctx context.Context, req *pb.GetArchitectureRequest) (*pb.GetArchitectureResponse, error) {
	return &pb.GetArchitectureResponse{ArchitectureType: []pb.ArchitectureType{
		pb.ArchitectureType_AWS,
		pb.ArchitectureType_AZURE,
		pb.ArchitectureType_DIGITAL_OCEAN,
		pb.ArchitectureType_GCP,
	}}, nil
}
