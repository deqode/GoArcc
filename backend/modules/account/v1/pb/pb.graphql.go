// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	gql_ptypes_empty "github.com/ysugimoto/grpc-graphql-gateway/ptypes/empty"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_UpdateAccountRequest     *graphql.Object      // message UpdateAccountRequest in account.proto
	gql__type_GetUserAccountsResponse  *graphql.Object      // message GetUserAccountsResponse in account.proto
	gql__type_GetUserAccountsRequest   *graphql.Object      // message GetUserAccountsRequest in account.proto
	gql__type_GetAccountRequest        *graphql.Object      // message GetAccountRequest in account.proto
	gql__type_DeleteAccountRequest     *graphql.Object      // message DeleteAccountRequest in account.proto
	gql__type_Account                  *graphql.Object      // message Account in account.proto
	gql__input_UpdateAccountRequest    *graphql.InputObject // message UpdateAccountRequest in account.proto
	gql__input_GetUserAccountsResponse *graphql.InputObject // message GetUserAccountsResponse in account.proto
	gql__input_GetUserAccountsRequest  *graphql.InputObject // message GetUserAccountsRequest in account.proto
	gql__input_GetAccountRequest       *graphql.InputObject // message GetAccountRequest in account.proto
	gql__input_DeleteAccountRequest    *graphql.InputObject // message DeleteAccountRequest in account.proto
	gql__input_Account                 *graphql.InputObject // message Account in account.proto
)

func Gql__type_UpdateAccountRequest() *graphql.Object {
	if gql__type_UpdateAccountRequest == nil {
		gql__type_UpdateAccountRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UpdateAccountRequest",
			Fields: graphql.Fields{
				"account": &graphql.Field{
					Type: Gql__type_Account(),
				},
			},
		})
	}
	return gql__type_UpdateAccountRequest
}

func Gql__type_GetUserAccountsResponse() *graphql.Object {
	if gql__type_GetUserAccountsResponse == nil {
		gql__type_GetUserAccountsResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserAccountsResponse",
			Fields: graphql.Fields{
				"accounts": &graphql.Field{
					Type: graphql.NewList(Gql__type_Account()),
				},
			},
		})
	}
	return gql__type_GetUserAccountsResponse
}

func Gql__type_GetUserAccountsRequest() *graphql.Object {
	if gql__type_GetUserAccountsRequest == nil {
		gql__type_GetUserAccountsRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserAccountsRequest",
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetUserAccountsRequest
}

func Gql__type_GetAccountRequest() *graphql.Object {
	if gql__type_GetAccountRequest == nil {
		gql__type_GetAccountRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAccountRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetAccountRequest
}

func Gql__type_DeleteAccountRequest() *graphql.Object {
	if gql__type_DeleteAccountRequest == nil {
		gql__type_DeleteAccountRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_DeleteAccountRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_DeleteAccountRequest
}

func Gql__type_Account() *graphql.Object {
	if gql__type_Account == nil {
		gql__type_Account = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_Account",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.String,
					Description: `id generated by uuid`,
				},
				"slug": &graphql.Field{
					Type:        graphql.String,
					Description: `multiple account identify by slug on ui`,
				},
				"user_id": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `userId of a user`,
				},
			},
		})
	}
	return gql__type_Account
}

func Gql__input_UpdateAccountRequest() *graphql.InputObject {
	if gql__input_UpdateAccountRequest == nil {
		gql__input_UpdateAccountRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UpdateAccountRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"account": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Account(),
				},
			},
		})
	}
	return gql__input_UpdateAccountRequest
}

func Gql__input_GetUserAccountsResponse() *graphql.InputObject {
	if gql__input_GetUserAccountsResponse == nil {
		gql__input_GetUserAccountsResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserAccountsResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"accounts": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_Account()),
				},
			},
		})
	}
	return gql__input_GetUserAccountsResponse
}

func Gql__input_GetUserAccountsRequest() *graphql.InputObject {
	if gql__input_GetUserAccountsRequest == nil {
		gql__input_GetUserAccountsRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserAccountsRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetUserAccountsRequest
}

func Gql__input_GetAccountRequest() *graphql.InputObject {
	if gql__input_GetAccountRequest == nil {
		gql__input_GetAccountRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAccountRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetAccountRequest
}

func Gql__input_DeleteAccountRequest() *graphql.InputObject {
	if gql__input_DeleteAccountRequest == nil {
		gql__input_DeleteAccountRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_DeleteAccountRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_DeleteAccountRequest
}

func Gql__input_Account() *graphql.InputObject {
	if gql__input_Account == nil {
		gql__input_Account = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_Account",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Description: `id generated by uuid`,
					Type:        graphql.String,
				},
				"slug": &graphql.InputObjectFieldConfig{
					Description: `multiple account identify by slug on ui`,
					Type:        graphql.String,
				},
				"user_id": &graphql.InputObjectFieldConfig{
					Description: `userId of a user`,
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_Account
}

// graphql__resolver_Accounts is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_Accounts struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_Accounts creates pointer of service struct
func new_graphql_resolver_Accounts(conn *grpc.ClientConn) *graphql__resolver_Accounts {
	return &graphql__resolver_Accounts{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_Accounts) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_Accounts) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"account": &graphql.Field{
			Type: Gql__type_Account(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetAccountRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for account")
				}
				client := NewAccountsClient(conn)
				resp, err := client.GetAccount(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetAccount")
				}
				return resp, nil
			},
		},
		"userAccounts": &graphql.Field{
			Type: Gql__type_GetUserAccountsResponse(),
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetUserAccountsRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for userAccounts")
				}
				client := NewAccountsClient(conn)
				resp, err := client.GetUserAccounts(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUserAccounts")
				}
				return resp, nil
			},
		},
		"deleteAccount": &graphql.Field{
			Type: gql_ptypes_empty.Gql__type_Empty(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req DeleteAccountRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for deleteAccount")
				}
				client := NewAccountsClient(conn)
				resp, err := client.DeleteAccount(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC DeleteAccount")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_Accounts) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"updateAccount": &graphql.Field{
			Type: Gql__type_Account(),
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{
					Type: Gql__input_Account(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req UpdateAccountRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for updateAccount")
				}
				client := NewAccountsClient(conn)
				resp, err := client.UpdateAccount(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC UpdateAccount")
				}
				return resp, nil
			},
		},
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterAccountsGraphqlHandler with *grpc.ClientConn manually.
func RegisterAccountsGraphql(mux *runtime.ServeMux) error {
	return RegisterAccountsGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service Accounts {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterAccountsGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_Accounts(conn))
}
