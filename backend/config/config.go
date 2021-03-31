package config

import (
	"alfred/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

//Config : including all the configuration
type Config struct {
	Grpc        GraphqlServerConfig
	Graphql     GraphqlServerConfig
	Rest        RestServerConfig
	HealthCheck HealthCheckServerConfig
	Logger      LoggerConfig
	Postgres    PostgresConfig
	Metrics     MetricsConfig
	Jaeger      JaegerServerConfig
	Auth        AuthConfig
}

// GrpcServerConfig: gRPC  server configuration
type GrpcServerConfig struct {
	Port string
	Host string
}

// GraphqlServerConfig: Graphql server configuration
type GraphqlServerConfig struct {
	Port string
	Host string
}

//RestServerConfig: Rest Implementation config
type RestServerConfig struct {
	Port string
	Host string
}

//HealthCheckServerConfig: Configuration about health check
type HealthCheckServerConfig struct {
	Port string
	Host string
}

//LoggerConfig: Zapier log level
type LoggerConfig struct {
	LogLevel string
}

//PostgresConfig: detail config about the postgres database
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSslmode  string
	PgDriver           string
}

//MetricsConfig : detail config about the Metrics
type MetricsConfig struct {
	Url         string
	ServiceName string
}

//JaegerServerConfig : detail config about the Jaeger
type JaegerServerConfig struct {
	Host        string
	Port        string
	ServiceName string
	LogSpans    string
}

type AuthConfig struct {
	Auth0ClientId     string
	Auth0Domain       string
	Auth0ClientSecret string
	Auth0CallbackUrl  string
}

// LoadConfig config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig file from the given viper
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		logger.Log.Fatal("unable to decode into struct", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

// GetConfig: get the path from local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "config_docker"
	}
	return "config_local"
}

// GetConfig : will get the config
func GetConfig() *Config {
	configPath := GetConfigPath(os.Getenv("config"))
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		logger.Log.Fatal("unable to get config", zap.Error(err))
		return nil
	}
	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		logger.Log.Fatal("unable to get config", zap.Error(err))
		return nil
	}
	return cfg
}
