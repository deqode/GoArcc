package ArchitectureSuggesterService

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/modules/ArchitectureSuggesterService/pb"
)

type ArchitectureSuggesterService struct {
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewArchitectureSuggesterService todo : AlWays add migration code for best practices
func NewArchitectureSuggesterService(
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.ArchitectureSuggesterServiceServer {

	return &ArchitectureSuggesterService{
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
