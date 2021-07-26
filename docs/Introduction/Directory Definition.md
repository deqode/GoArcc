---
sidebar_position: 3
---




Before you start writing your services you need to understand how GoArcc works
and what is the definition of its directory. 
Let's get deep dive into the definition and codebase.

# <u>Directory Structure</u>



We have followed a standard go directory structure for our codebase so that developers will easily get friendly with it.
we have multiple directories inside GoArcc. we have used multiples of libraries, we will go through each one by one.

In below diagram we have multiple of directories such as **cmd** , **client** , **logger** , **protos** **etc**.

<!-- <img src="/img/folderDefinition/goArcc.jpg" alt="GoArcc Directory Structure" width="20%"/> -->

![GoArcc Directory Structure](/img/folderDefinition/goArcc.jpg)

## <u>Config</u>



When we start the GoArcc application we need some information so that our application will start smoothly. 
As we have said earlier, we are supporting gRPC , rest, graphql . So we need a different port
for all the servers and different settings to start the application.

We have config.yaml file in the root directory. 
You can change any information from the file like id, password of database, etc according to your use case.

see the below example of config.yaml file.



```yaml

grpc:
  port: 8000
  host: localhost
  request_timeout: 20

graphql:
  port: 8081
  host: localhost
  request_Timeout: 20

rest:
  port: 8082
  host: localhost
  request_timeout: 20

```

### <u>How you can add custom config?</u>


If you want to add some extra configs then you need to make a struct and place inside the main struct.


```go

type Config struct {
	Grpc               GrpcServerConfig        `mapstructure:"GRPC"`
	Graphql            GraphqlServerConfig     `mapstructure:"GRAPHQL"`
	Rest               RestServerConfig        `mapstructure:"REST"`
	HealthCheck        HealthCheckServerConfig `mapstructure:"HEALTH_CHECK"`
	Logger             LoggerConfig            `mapstructure:"LOGGER"`
	Postgres           PostgresConfig          `mapstructure:"POSTGRES"`
	Metrics            MetricsConfig           `mapstructure:"METRICS"`
	Jaeger             JaegerServerConfig      `mapstructure:"JAEGER"`
	Auth               AuthConfig              `mapstructure:"AUTH"`
	GithubVCSConfig    VCSSConfig              `mapstructure:"GITHUB_VCS_CONFIG"`
	CadenceConfig      CadenceConfig           `mapstructure:"CADENCE_CONFIG"`
}

```



:::tip Bright Side

If you pass the environment variables then those variables will be overridden
to the default variables which is present in the config.yaml file.

:::

### <u>How you can pass enviroment variables?</u>

```

GRPC_PORT = 8085

```


## <u>Db</u>



Inside connection.go file we have a method called NewConnection().
NewConnection is responsible for open the database connection. 
Db instance is bind with the fx provider so we don't need to call this method.

See the below code.

```go

func NewConnection(config *config.Config) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable Timezone=Asia/Shanghai",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName)

	// Refrence taken from https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Log.Fatal("GORM connection failed", zap.Error(err))
		panic(err)
	}

	logger.Log.Info("connection established with the database")
	//No need to close the connection because we have a single pool of connection.
	return db
}

```


In the above code you can change the setting like timezone , ssl mode etc.



:::tip Bright Side

We have used GORM as an ORM. If you want to explore GORM then once go through the docs.
(https://gorm.io/docs/index.html)

:::



## <u>Logger</u>



As we know logging is an essential part of the application. 
Logging helps in troubleshooting the application and you can also see the performance of your infrastructure.
Logging provides more visibility to your application at the components level. 
Logs contain important information so that anyone can debug or find the fault if exist.

We have implemented log in the whole application so that developers will no need to worry about logging.

In golang, we have multiple logging libraries. So the team has decided to use a zap logger.


### <u>Why we have used zap logger? </u>

Just see the comparison below 😀

Log a message and 10 fields:

| Package | Time | Time % to zap | Objects Allocated |
| :------ | :--: | :-----------: | :---------------: |
| :zap: zap | 862 ns/op | +0% | 5 allocs/op
| :zap: zap (sugared) | 1250 ns/op | +45% | 11 allocs/op
| zerolog | 4021 ns/op | +366% | 76 allocs/op
| go-kit | 4542 ns/op | +427% | 105 allocs/op
| apex/log | 26785 ns/op | +3007% | 115 allocs/op
| logrus | 29501 ns/op | +3322% | 125 allocs/op
| log15 | 29906 ns/op | +3369% | 122 allocs/op




In main method there is a method called logger.Init().
This method will create a global logger for your application and takes log level from config.yaml
file.

Apart from that if we need to overide the log level we just set an env variable or you can also change in the 
config.yaml file.

### <u>Zap log level reference </u>


```go
// reference taken from https://github.com/uber-go/zap/blob/120b08c8fae5be92bc7323ca78b25cd790e8c37a/level.go#L28
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

```



:::tip Bright Side

Zap logger have different log level. You can read more about zap (https://pkg.go.dev/go.uber.org/zap)
You can also see the different log level of zap
:::



## <u>Cmd</u>




As we know in golang we have a convention of cmd folder 
and it has its own significance.
If you want to read about the go project layout then go through 
this doc once (https://github.com/golang-standards/project-layout).

lets come to the our code base.
when we go inside the cmd folder we see there are multiples of files(paste github link here).
 
 - **main.go**
 - **invokers.go**
 - **providers.go**
 - **app.go**

:::tip Bright Side

Before understanding the use of each file first go through dependency
management in golang.
(https://github.com/uber-go/fx)

:::

   ### <u>main.go</u>
     

```go

package main

import (
	"go.uber.org/zap"
	"goarcc/logger"
)

func main() {
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, 
		Development: false,
	})
	GetApp().Run()
}

```



As we know main.go is the entry point for any go application.
Here logger.Init() will take the config object as a parameter and initialize the whole application with debug level.
This means in the standard output you will able to see the debug levels of the log only.



:::tip Tip

If you want to see the info or warn level logs then you need to just pass the Zap loglevel in config object.

:::


   ### <u>invokers.go</u>
    
```go

package main

import (
	"go.uber.org/fx"
	"goarcc/servers/cleanup"
	"goarcc/servers/graphql"
	"goarcc/servers/grpc"
	"goarcc/servers/healthcheck"
	"goarcc/servers/rest"
)


func GetInvokersOptions() fx.Option {
	return fx.Invoke(
		grpc.RunGRPCServer,
		grpc.RegisterGrpcModules,
		rest.RunRestServer,
		graphql.RunGraphqlServer,
		healthcheck.HealthCheckRunner,
		cleanup.Cleanup,
	)
}

```



Let's understand the above code. Here We have a method GetInvokersOptions(). 
This function is responsible to start multiple servers.
we have gRPC , rest, graphql server. Apart from that we also have a health check server.
cleanup code will be executed when we stop our application. 
Generally, In cleanup code, we gracefully shutdown everything and clears the cache
if any present.



:::danger Alarm

Inside Invoker, we have multiple methods. The invoker method will execute each method sequentially.
So there are multiple dependencies among them and if you change the order of execution you will start getting lots of errors.
So, if you have explored the fx library then you can play with invokers and providers.

:::

### <u>providers.go</u>

```go

package main

import (
	"go.uber.org/fx"
	"goarcc/client/grpcClient"
	"goarcc/config"
	"goarcc/db"
	"goarcc/servers/cleanup"
	"goarcc/servers/grpc"
	"goarcc/servers/openTracing/tracer/jaeger"
)

func GetProviderOptions() []fx.Option {
	return []fx.Option{
		config.ProviderFx,
		grpc.InitGrpcBeforeServingFx,
		db.DatabaseConnectionFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
		grpcClient.ConnectionFx,
	}
}

```



Let's understand the above code. Here We have a method GetProvidersOptions(). 
This function load all the dependency which will need to start the application.
There are multiples of server and each server have some sort of dependency or need some predefined object
at runtime.



### <u>app.go</u>

```go

package main

import "go.uber.org/fx"

func GetApp() *fx.App {
	opts := make([]fx.Option, 0)
	opts = GetProviderOptions()
	opts = append(opts, GetInvokersOptions())
	return fx.New(
		opts...,
	)
}

```

Here GetApp() method will resolve  all the dependency and bind the whole application as an object. 
So , In main file you can easily start the application.



## <u> Servers </u>



<!-- <img src="/img/folderDefinition/servers.jpg" alt="GoArcc Servers Structure" width="70%"/>-->


![GoArcc Servers Structure](/img/folderDefinition/Servers.jpg)

Servers are the core part of this project. you can see in the above that we have supported multiples of servers.
lets go into the deep dive of the code base so you can understand easily and change accordingly.

### GRPC
GoArcc built upon client server architecture. So GRPC is core server for GoArcc, every request is handled by grpc server and middlewares.
#### GRPC Middleware
You just have to setup your middleware into gRPC and it will work for rest and graphQL as well. Auth middleware already implemented here.
#### how to register your service client?
you can register your service as grpc at `servers/grpc/register.go`
```go

func RegisterGrpcModules(srv *grpc.Server,
	db *gorm.DB,
	config *config.Config,
	grpcClientConnection *grpc.ClientConn) {

	//register new grpc modules here
	// RegisterUserProfilesServer autogenerated in pb/..._grpc.go file
	userProfilePb.RegisterUserProfilesServer(srv, userExt.NewUserProfilesServer(db, config, grpcClientConnection))
}

```

### REST
REST API's over grpc is achieved by [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway). it basically translates a RESTful HTTP API into gRPC.
#### how to register your service as REST?
you can register your service as rest at `servers/rest/register.go`
```go
func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
   // Register your REST module here
   // RegisterUserProfilesHandler autogenerated in pb/..._gw.go
	if err := userProfilePb.RegisterUserProfilesHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
```

### GraphQL
We are using [grpc-graphql-gateway](https://github.com/ysugimoto/grpc-graphql-gateway), this plugin basically generate GraphQL Schema from Protocol Buffers.
#### how to register your service as graphQL?
you can register your service as rest at `servers/graphql/register.go`
```go
//RegisterGraphqlModules: Mapping the services with the single graphql endpoint
func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
    // Register your graphql service here
	if err := userProfilePb.RegisterUserProfilesGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
```


### HealthCheck
Any client will be able to check the health of our product. Health Information gives a brief idea to any user that the service of your project is currently down.
HealthCheck is running on `http://localhost:8083/health/`

### Jaeger
Jaeger is open source software for tracing transactions between distributed services. It’s used for monitoring and troubleshooting complex microservices environments.
Open `http://loaclhost:16686`.
Now you will be able to trace each request easily. [learn more](https://e)

### Prometheus
Prometheus is an open-source systems monitoring and alerting.
Open `http://localhost:9090`.


