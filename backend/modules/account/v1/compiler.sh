#!/usr/bin/env sh
protoc -I .  \
       --proto_path=../../../protos  \
       --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true \
       --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false \
       --graphql_out=./  \
       --validate_out="lang=go:./"  ./*.proto
go fmt ./...
go mod tidy
