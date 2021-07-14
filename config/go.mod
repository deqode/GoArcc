module github.com/deqode/GoArcc/config

go 1.16

require (
	github.com/spf13/viper v1.8.1
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.18.1
	github.com/deqode/GoArcc/logger v0.0.0-00010101000000-000000000000
)

replace github.com/deqode/GoArcc/logger => ./../logger
