package external_svc

import (
	"alfred/config"
	model "alfred/modules/stack/v1/models"
	"alfred/modules/stack/v1/pb"
	"gorm.io/gorm"
)

type stacksServer struct {
	db     *gorm.DB
	config *config.Config
}

// NewStacksServer todo : AlWays add migration code for best practices
func NewStacksServer(
	db *gorm.DB,
	config *config.Config,
) pb.StacksServer {

	//initial migration of databases: schema migration
	model.InitialMigrationStack(db)
	return &stacksServer{
		db:     db,
		config: config,
	}
}
