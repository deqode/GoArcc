#!/usr/bin/env sh
# put the protoc path of your local
protoc -I .  --proto_path=/home/atul/Desktop/deqodeEnviroment/protos  --grpc-gateway_out .     --grpc-gateway_opt logtostderr=true    --grpc-gateway_opt paths=source_relative    ./pb/*.proto
protoc -I .  --proto_path=/home/atul/Desktop/deqodeEnviroment/protos   --go_out=./pb   --graphql_out=./pb --go-grpc_out=./pb ./pb/*.proto
go fmt ./...
go mod tidy
