#!/usr/bin/env sh
# put the protoc path of your local
protoc -I .  --proto_path=../../../../protos  --grpc-gateway_out ./pb     --grpc-gateway_opt logtostderr=true    --grpc-gateway_opt paths=source_relative    ./*.proto
protoc -I .  --proto_path=../../../../protos    --go_out=./pb    --go-grpc_out=./pb ./*.proto
go fmt ./...
go mod tidy
