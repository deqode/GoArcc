module goarcc

go 1.12

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dimiro1/health v0.0.0-20191019130555-c5cbb4d46ffc
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.5.2
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/graphql-go/graphql v0.7.9
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/labstack/echo/v4 v4.2.1
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.5
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.0.0-20201205024021-ac21108117ac // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/smartystreets/assertions v1.1.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	github.com/ysugimoto/grpc-graphql-gateway v0.20.0
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.18.1
	goarcc/config v0.0.0-00010101000000-000000000000
	goarcc/db v0.0.0-00010101000000-000000000000
	goarcc/logger v0.0.0-00010101000000-000000000000
	goarcc/protos/types v0.0.0-00010101000000-000000000000
	goarcc/servers/grpc v0.0.0-00010101000000-000000000000
	goarcc/util/userinfo v0.0.0-00010101000000-000000000000
	golang.org/x/exp v0.0.0-20200331195152-e8c3332aa8e5 // indirect
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.11
)

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02
	goarcc/config => ./config
	goarcc/db => ./db
	goarcc/logger => ./logger
	goarcc/protos/types => ./protos/types
	goarcc/servers/grpc => ./servers/grpc
	goarcc/util/userinfo => ./util/userinfo
)
