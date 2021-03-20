package aws

import (
	"alfred/logger"
	"alfred/modules/ArchitectureSuggesterService/aws/awsDefault"
	awsDefaultPb "alfred/modules/ArchitectureSuggesterService/aws/awsDefault/pb"
	"alfred/modules/ArchitectureSuggesterService/aws/pb"
	"go.uber.org/zap"
)

type AwsService struct {
	awsDefaultData *awsDefaultPb.AwsDefaultData
}


func NewAwsService() pb.AwsServiceServer {
	data, err := awsDefault.ParserData("awsDefault.yml")
	if err != nil {
		logger.Log.Fatal("not able to register", zap.Error(err))
	}
	return &AwsService{
		awsDefaultData: data,
	}
}
