package aws

import (
	awsdefaultPb "alfred/modules/ArchitectureSuggesterService/aws/awsDefault/pb"
	"alfred/modules/ArchitectureSuggesterService/aws/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a AwsService) GetAwsMachineTypes(ctx context.Context, request *pb.GetAwsMachineTypesRequest) (*pb.GetAwsMachineTypesResponse, error) {

	if a.awsDefaultData == nil ||
		a.awsDefaultData.AwsContainerServices == nil {
		return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting aws container service")
	}

	machineTypes := make([]string, 0)
	machineConfig := make([]*pb.MachineConfig, 0)

	switch request.Parent {
	case "ECS":
		if a.awsDefaultData.AwsContainerServices.EcsContainerService == nil {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting ecs container service")
		}
		if a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType == nil {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting ecs  launch type")
		}

		switch request.InstanceType {
		case "GENERAL_PURPOSE":
			if a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.GeneralPurposeMachine != nil {
				return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting general purpose machine")
			}
			machineTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.GeneralPurposeMachine.Name
			machineConfig = configTypeConverter(a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.GeneralPurposeMachine.MachineConfig)
		case "COMPUTE_OPTIMIZED":
			if a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.ComputeOptimizedMachine != nil {
				return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting compute optimized")
			}
			machineTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.ComputeOptimizedMachine.Name
			machineConfig = configTypeConverter(a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.ComputeOptimizedMachine.MachineConfig)

		case "MEMORY_OPTIMIZED":
			if a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.MemoryOptimizedMachine != nil {
				return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting memory optimized")
			}
			machineTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.MemoryOptimizedMachine.Name
			machineConfig = configTypeConverter(a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.MemoryOptimizedMachine.MachineConfig)

		case "ACCELERATED_COMPUTING":
			if a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.AccelerateComputing != nil {
				return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting accelerated computing")
			}
			machineTypes = a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.AccelerateComputing.Name
			machineConfig = configTypeConverter(a.awsDefaultData.AwsContainerServices.EcsContainerService.Ec2LaunchType.AccelerateComputing.MachineConfig)

		default:
			return nil, status.Error(codes.InvalidArgument, "invalid launch type or launch type not exist")
		}

	case "EKS":
		if a.awsDefaultData.AwsContainerServices.EksContainerService == nil {
			return nil, status.Error(codes.FailedPrecondition, "something went wrong while getting eks container service")
		}
		//todo will implement it later

	}

	return &pb.GetAwsMachineTypesResponse{MachineTypes: machineTypes, MachineConfig: machineConfig}, nil
}






func configTypeConverter(configs []*awsdefaultPb.MachineConfig) []*pb.MachineConfig {
	machineConfig := make([]*pb.MachineConfig, 0)
	for _, config := range configs {
		machineConfig = append(machineConfig, &pb.MachineConfig{
			Parent:                    config.Parent,
			InstanceSize:              config.InstanceSize,
			VCpu:                      config.VCpu,
			MemoryInGib:               config.MemoryInGib,
			InstanceStorage:           config.InstanceStorage,
			InstanceStorageGib:        config.InstanceStorageGib,
			NetworkBandwidthGbpc:      config.NetworkBandwidth,
			EbsBandwidthMbps:          config.EbsBandwidth,
			BaselinePerformance:       config.BaselinePerformance,
			CpuCreditsEarnedHr:        config.CpuCreditsEarnedHr,
			NetworkBurstBandwidthGbps: config.NetworkBandwidthGbpc,
			EbsBurstBandwidthMbps:     config.EbsBandwidthMbps,
			Mem:                       config.Mem,
			MemGib:                    config.MemoryInGib,
			NetworkPerformance:        config.NetworkPerformance,
			Storage:                   config.Storage,
			StorageGb:                 config.StorageGb,
			NetworkPerformanceGib:     config.NetworkPerformanceGib,
			NetworkBandwidth:          config.NetworkBandwidth,
			EbsBandwidth:              config.EbsBandwidth,
			GpuMemGib:                 config.GpuMemGib,
			GpuMem:                    config.GpuMem,
			GpuP2P:                    config.GpuP2P,
			DedicatedEbsBandwidth:     config.DedicatedEbsBandwidth,
			LocalStorage:              config.LocalStorage,
			LocalStorageGb:            config.LocalStorageGb,
		})
	}
	return machineConfig
}
