package ArchitectureSuggesterService

import (
	"alfred/modules/ArchitectureSuggesterService/pb"
)

type ArchitectureSuggesterService struct {
}

//todo : AlWays add migration code for best practices
func NewArchitectureSuggesterService() pb.ArchitectureSuggesterServiceServer {

	return &ArchitectureSuggesterService{}
}
