package db

import (
	"alfred/config"
	"alfred/logger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection NewConnection: will open the connection with the database information
// that is passed as an argument
func NewConnection(config *config.Config) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable Timezone=Asia/Shanghai",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName)

	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Log.Fatal("GORM connection failed", zap.Error(err))
		panic(err)
	}

	logger.Log.Info("connection established with the database")
	//No need to close the connection because we have a single pool of connection
	return db
}
