#!/usr/bin/env sh
protoc -I .  \
       --proto_path=../../../../protos  \
       --proto_path=../ \
       --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true \
       --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false \
       --graphql_out=./  \
       ./*.proto
go fmt ./...
go mod tidy
