#!/usr/bin/env sh
# put the protoc path of your local
protoc -I .  --proto_path=../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false --graphql_out=./  ./*.proto
#protoc -I .   -I ${GOPATH}/src -I /home/shivang/go/src/github.com/envoyproxy/protoc-gen-validate/  --go_out=./pb  --validate_out="lang=go:./pb" ./*.proto
go fmt ./...
go mod tidy
