package background

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type CadenceConfig struct {
	Domain   string
	Service  string
	HostPort string
}

type CadenceAppConfig struct {
	Env            string
	WorkerTaskList string
	Cadence        CadenceConfig
	Logger         *zap.Logger
}

// Setup setup the config for the code run
func (h *CadenceAppConfig) Setup() {
	viper.SetConfigName("background_local")
	viper.AddConfigPath("background") // These two lines will make sure viper pulls the config from app/resources/application.yml
	viper.AutomaticEnv()                 // This allows viper to read variables from the environment variables if they exists.
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&h)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	h.Logger = logger

	logger.Debug("Finished loading Configuration!")
}

