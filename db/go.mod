module github.com/deqode/GoArcc/db

go 1.16

require (
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.18.1
	github.com/deqode/GoArcc/config v0.0.0-00010101000000-000000000000
	github.com/deqode/GoArcc/logger v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)

replace (
	github.com/deqode/GoArcc/config => ./../config
	github.com/deqode/GoArcc/logger => ./../logger
)
