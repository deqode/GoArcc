package user_profile

// - protoc plugins must be in $PATH
//go:generate bash -e -o pipefail -c "protoc -I .  --proto_path=../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false --graphql_out=./  --validate_out='lang=go:./'  ./pb/*.proto"
