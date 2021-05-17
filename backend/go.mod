module alfred

go 1.16

require (
	alfred.sh/common v1.0.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimiro1/health v0.0.0-20191019130555-c5cbb4d46ffc
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/go-redis/redis/v8 v8.8.2
	github.com/gogo/googleapis v1.3.2 // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v34 v34.0.0
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/graphql-go/graphql v0.7.9
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/justinas/alice v1.2.0
	github.com/kisielk/errcheck v1.6.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.2.2
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus/client_golang v1.3.0
	github.com/samuel/go-thrift v0.0.0-20191111193933-5165175b40af // indirect
	github.com/smartystreets/assertions v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/uber-go/tally v3.3.17+incompatible
	github.com/uber/jaeger-client-go v2.27.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/uber/tchannel-go v1.16.0 // indirect
	github.com/ysugimoto/grpc-graphql-gateway v0.20.0
	go.uber.org/cadence v0.17.0
	go.uber.org/fx v1.13.1
	go.uber.org/net/metrics v1.3.0 // indirect
	go.uber.org/thriftrw v1.25.0 // indirect
	go.uber.org/yarpc v1.42.1
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/tools v0.1.1 // indirect
	google.golang.org/genproto v0.0.0-20210427215850-f767ed18ee4d
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02

replace alfred.sh/common => ./../common
