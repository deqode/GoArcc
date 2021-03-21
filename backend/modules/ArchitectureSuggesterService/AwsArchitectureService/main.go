package AwsArchitectureService

import (
	"alfred/logger"
	"alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/ArchitecureConfiguration"
	awsDefaultPb "alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/ArchitecureConfiguration/pb"
	"alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/pb"
	"go.uber.org/zap"
)

type AwsArchitectureService struct {
	awsDefaultData *awsDefaultPb.AwsDefaultData
}

func NewAwsArchitectureService() pb.AwsServiceServer {
	data, err := ArchitecureConfiguration.ParserData("awsConfig.yml")
	if err != nil {
		logger.Log.Fatal("not able to register", zap.Error(err))
	}
	return &AwsArchitectureService{
		awsDefaultData: data,
	}
}
