package user_profile

// generate from pb/user_profile.proto
//go:generate bash -e -o pipefail -c "protoc -I .  --proto_path=../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false --graphql_out=./  --validate_out='lang=go:./'  ./pb/user_profile.proto"

// generate from pb/user_profile_int.proto
//go:generate bash -e -o pipefail -c "protoc -I .  --proto_path=../../../protos --proto_path=./pb  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --go-grpc_opt=require_unimplemented_servers=false  --validate_out='lang=go:./'  ./pb/user_profile_int.proto"
