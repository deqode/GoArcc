package config

import (
	"alfred/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

/*// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	//Graphql Port
	GraphqlPort string

	//Promthesius Port
	PromthesiusPort string

	ZipkinUrl string

	ServerHost string
	//
	DatastoreDBPort int

	ZipkinServiceName string

	HealthCheckPort string
	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

*/

type Config struct {
	Grpc        GraphqlServerConfig
	Graphql     GraphqlServerConfig
	Rest        RestServerConfig
	HealthCheck HealthCheckServerConfig
	Promthesius PromthesiusServerConfig
	Logger      LoggerConfig
	Postgres    PostgresConfig
	Metrics     MetricsConfig
	Jaeger      JaegerServerConfig
}

type GrpcServerConfig struct {
	Port string
	Host string
}

type GraphqlServerConfig struct {
	Port string
	Host string
}

type RestServerConfig struct {
	Port string
	Host string
}

type HealthCheckServerConfig struct {
	Port string
	Host string
}

type PromthesiusServerConfig struct {
	Port string
	Host string
}

type LoggerConfig struct {
	LogLevel string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSslmode  string
	PgDriver           string
}

type MetricsConfig struct {
	Url         string
	ServiceName string
}
type JaegerServerConfig struct {
	Host        string
	Port        string
	ServiceName string
	LogSpans    string
}

/*func GetConfig() *Config {
	return &Config{
		GraphqlPort:         "8082",
		HTTPPort:            "8081",
		GRPCPort:            "8080",
		ServerHost:          "localhost",
		LogLevel:            -1,
		DatastoreDBPort:     5432,
		DatastoreDBUser:     "alfred",
		DatastoreDBPassword: "alfred",
		DatastoreDBSchema:   "alfred.v1",
		DatastoreDBHost:     "localhost",
		PromthesiusPort:     "8083",
		HealthCheckPort:     "8084",
		ZipkinUrl:           "http://localhost:9411/api/v1/spans",
		ZipkinServiceName:   "AlfredTracing",
	}
}

*/

// Load config file from given path
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

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		logger.Log.Fatal("unable to decode into struct", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "config-docker"
	}
	return "config_local"
}

// Get config
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
