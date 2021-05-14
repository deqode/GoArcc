package config

import (
	"alfred/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

type FileInformation struct {
	Path string
	Name string
}

//Config : including all the configuration
type Config struct {
	Grpc               GrpcServerConfig
	Graphql            GraphqlServerConfig
	Rest               RestServerConfig
	HealthCheck        HealthCheckServerConfig
	Logger             LoggerConfig
	Postgres           PostgresConfig
	Metrics            MetricsConfig
	Jaeger             JaegerServerConfig
	Auth               AuthConfig
	GithubVCSConfig    VCSSConfig
	SupportedVcsConfig []string `ignored:"true"`
	ConfigLocal        bool
}

// GrpcServerConfig GrpcServerConfig: gRPC  server configuration
// Timeout is the request timeout:
// any client request take longer then the given timeout will automatically cancelled.
type GrpcServerConfig struct {
	Port           string
	Host           string
	RequestTimeout int
}

// GraphqlServerConfig GraphqlServerConfig: Graphql server configuration
type GraphqlServerConfig struct {
	Port           string
	Host           string
	RequestTimeout int
}

// RestServerConfig RestServerConfig: Rest Implementation config
type RestServerConfig struct {
	Port           string
	Host           string
	RequestTimeout int
}

// HealthCheckServerConfig HealthCheckServerConfig: Configuration about health check
type HealthCheckServerConfig struct {
	Port string
	Host string
}

// LoggerConfig LoggerConfig: Zapier log level
type LoggerConfig struct {
	LogLevel string
}

// PostgresConfig PostgresConfig: detail config about the postgres database
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
	URL         string
	ServiceName string
}

//JaegerServerConfig : detail config about the Jaeger
type JaegerServerConfig struct {
	Host        string
	Port        string
	ServiceName string
	LogSpans    string
}

// AuthConfig Authentication config: details provided by Auth0
type AuthConfig struct {
	Auth0ClientID     string
	Auth0Domain       string
	Auth0ClientSecret string
	Auth0CallbackURL  string
}

//VCSSConfig like github,gitlab,bitbucket config
type VCSSConfig struct {
	IType        int
	Provider     string
	URLTemplate  string
	ClientID     string
	RedirectURI  string
	State        string
	Scope        string
	ResponseType string
	ClientSecret string
	Name         string
}

//GetVcsConfig : will give the particular vcs config
func GetVcsConfig(name string, vcsConfig []VCSSConfig) *VCSSConfig {
	for _, v := range vcsConfig {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

// LoadConfig config file from given path
func LoadConfig(filename, path string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(path)
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

// GetConfigPath get the path from local or docker
func GetConfigPath(configFileName string) string {
	if configFileName == "docker" {
		return "config_docker"
	}
	return "config_local"
}

// GetConfig : will get the config
func GetConfig(fileInformation *FileInformation) *Config {
	var config Config
	err := envconfig.Process("API", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.ConfigLocal {
		configFileName := GetConfigPath(os.Getenv("config"))
		cfgFile, err := LoadConfig(configFileName, ".")
		if err != nil {
			logger.Log.Fatal("unable to get config", zap.Error(err))
			return nil
		}
		cfg, err := ParseConfig(cfgFile)
		if err != nil {
			logger.Log.Fatal("unable to get config", zap.Error(err))
			return nil
		}
		cfg.SupportedVcsConfig = supportedVcsConfig()
		return cfg
	}
	config.SupportedVcsConfig = supportedVcsConfig()
	return &config
}

//GetFileInformation : will get the information of file
func GetFileInformation() *FileInformation {
	return &FileInformation{
		Name: "config_local",
		Path: ".",
	}
}

// SupportedVcsConfig todo
// SupportedVcsConfig add supported type from here.
func supportedVcsConfig() []string {
	return []string{"github"}
}
