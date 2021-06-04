module alfred

go 1.12

require (
	alfred.sh/common v0.0.0-00010101000000-000000000000
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimiro1/health v0.0.0-20191019130555-c5cbb4d46ffc
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/go-redis/redis/v8 v8.7.1
	github.com/golang/protobuf v1.4.3
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v34 v34.0.0
	github.com/google/uuid v1.1.2
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/graphql-go/graphql v0.7.8
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/hashicorp/go-uuid v1.0.1
	github.com/justinas/alice v1.2.0
	github.com/kyoh86/richgo v0.3.7 // indirect
	github.com/kyoh86/xdg v1.2.0 // indirect
	github.com/labstack/echo/v4 v4.2.1
	github.com/lib/pq v1.9.0 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.5
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.0.0-20201205024021-ac21108117ac // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/smartystreets/assertions v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/uber-go/tally v3.3.15+incompatible
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	github.com/wacul/ptr v1.0.0 // indirect
	github.com/ysugimoto/grpc-graphql-gateway v0.20.0
	go.uber.org/cadence v0.16.1-0.20210420001847-b31c3bbfb06a
	go.uber.org/fx v1.13.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/yarpc v1.53.2
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20210224155714-063164c882e6
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20210514222045-a12250e98f97 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.10
)

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02

replace alfred.sh/common => ./../common
