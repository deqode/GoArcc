module goarcc/db

go 1.16

require (
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.18.1
	goarcc/config v0.0.0-00010101000000-000000000000
	goarcc/logger v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)

replace (
	goarcc/config => ./../config
	goarcc/logger => ./../logger
)
