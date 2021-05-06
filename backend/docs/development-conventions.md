
``
rpc GetUserAllAccount(GetUserAllAccountRequest) returns (GetUserAllAccountResponse){
option (google.api.http) = {
get: "/v1/account/get-user-all-account/{user_id}"
};
option (graphql.schema) = {
type: QUERY
name: "userAllAccount"
};
}
``
