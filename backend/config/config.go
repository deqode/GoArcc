package config

// Config is configuration for Server
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

func GetConfig() *Config {
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
