package stack

// generate from pb/stack.proto
//go:generate bash -e -o pipefail -c "protoc -I .  --proto_path=../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false --graphql_out=./  --validate_out='lang=go:./'  ./pb/stack.proto"

// generate from pb/stack_build.proto
//go:generate bash -e -o pipefail -c "protoc -I .  --proto_path=../../../protos --proto_path=./pb  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false  --validate_out='lang=go:./'  ./pb/stack_build.proto"
