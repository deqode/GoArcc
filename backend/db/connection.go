package db

import (
	"alfred/config"
	"alfred/logger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//New Connection will open the connection with the database information
// that is passed as an argument.
func NewConnection(config *config.Config) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable Timezone=Asia/Shanghai",
		config.DatastoreDBHost, config.DatastoreDBPort, config.DatastoreDBUser, config.DatastoreDBPassword, config.DatastoreDBSchema)

	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("connection failed in db", zap.Error(err))
	}
	defer
	logger.Log.Info("connection established with the database")
	return db
}
