#!/usr/bin/env sh
# put the protoc path of your local
#protoc -I .  --proto_path=../../../protos  --grpc-gateway_out ./pb     --grpc-gateway_opt logtostderr=true    --grpc-gateway_opt paths=source_relative    ./*.proto
#protoc -I .  --proto_path=../../../protos    --go_out=/   --graphql_out=./pb --go-grpc_out=./pb ./*.proto
protoc -I .  --proto_path=../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false --graphql_out=./   ./*.proto
#protoc -I .  --proto_path=../../../protos     ./*.proto

go fmt ./...
go mod tidy

