package pb

// - protoc plugins must be in $PATH
//go:generate bash -e -o pipefail -c "protoc -I . --proto_path=../../../../protos --grpc-gateway_out=paths=source_relative:.  --grpc-gateway_opt logtostderr=true --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false --graphql_out=.   ./*.proto"
