module github.com/deqode/GoArcc/servers/grpc

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/justinas/alice v1.2.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.18.1
	github.com/deqode/GoArcc/config v0.0.0-00010101000000-000000000000
	github.com/deqode/GoArcc/logger v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
	gopkg.in/square/go-jose.v2 v2.6.0
	gorm.io/gorm v1.21.11
)

replace (
	github.com/deqode/GoArcc/config => ./../../config
	github.com/deqode/GoArcc/logger => ./../../logger
)
