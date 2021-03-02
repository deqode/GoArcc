package cmd

import (
	"alfred/config"
	grpql "alfred/protocol/graphql"
	"alfred/protocol/rest"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)



// RunServer runs gRPC server and HTTP gateway
func RunServer(lc fx.Lifecycle , cfg *config.Config)  {


	ctx := context.Background()
	// get configuration

	/*flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()*/

	cfg.GRPCPort = "8080"
	cfg.HTTPPort = "8081"
	cfg.GraphqlPort = "8082"
	cfg.LogLevel = int(zap.DebugLevel)

/*
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()*/

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// run HTTP gateway
			go func() {
				_ = rest.RunRESTServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
			}()
			//run graphql gateway
			go func() {
				_ = grpql.RunGraphqlServer(ctx , cfg.GraphqlPort)
			}()

			return nil
		},
	})

}

