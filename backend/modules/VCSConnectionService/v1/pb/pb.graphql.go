// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	types "alfred/protos/types"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"

	gql_ptypes_timestamp "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamp"
)

var (
	gql__type_VCSConnection                         *graphql.Object      // message VCSConnection in vcs-connection-service.proto
	gql__type_ListVCSConnectionResponse             *graphql.Object      // message ListVCSConnectionResponse in vcs-connection-service.proto
	gql__type_ListVCSConnectionRequest              *graphql.Object      // message ListVCSConnectionRequest in vcs-connection-service.proto
	gql__type_ListAllSupportedVCSProvidersResponse  *graphql.Object      // message ListAllSupportedVCSProvidersResponse in vcs-connection-service.proto
	gql__type_CallbackRequest                       *graphql.Object      // message CallbackRequest in vcs-connection-service.proto
	gql__type_AuthorizeResponse                     *graphql.Object      // message AuthorizeResponse in vcs-connection-service.proto
	gql__type_AuthorizeRequest                      *graphql.Object      // message AuthorizeRequest in vcs-connection-service.proto
	gql__type_AccountVCSConnection                  *graphql.Object      // message AccountVCSConnection in vcs-connection-service.proto
	gql__input_VCSConnection                        *graphql.InputObject // message VCSConnection in vcs-connection-service.proto
	gql__input_ListVCSConnectionResponse            *graphql.InputObject // message ListVCSConnectionResponse in vcs-connection-service.proto
	gql__input_ListVCSConnectionRequest             *graphql.InputObject // message ListVCSConnectionRequest in vcs-connection-service.proto
	gql__input_ListAllSupportedVCSProvidersResponse *graphql.InputObject // message ListAllSupportedVCSProvidersResponse in vcs-connection-service.proto
	gql__input_CallbackRequest                      *graphql.InputObject // message CallbackRequest in vcs-connection-service.proto
	gql__input_AuthorizeResponse                    *graphql.InputObject // message AuthorizeResponse in vcs-connection-service.proto
	gql__input_AuthorizeRequest                     *graphql.InputObject // message AuthorizeRequest in vcs-connection-service.proto
	gql__input_AccountVCSConnection                 *graphql.InputObject // message AccountVCSConnection in vcs-connection-service.proto
)

func Gql__type_VCSConnection() *graphql.Object {
	if gql__type_VCSConnection == nil {
		gql__type_VCSConnection = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_VCSConnection",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.String,
					Description: `id generated by uuid`,
				},
				"label": &graphql.Field{
					Type:        graphql.String,
					Description: `label of VCS Connection`,
				},
				"provider": &graphql.Field{
					Type:        types.Gql__enum_GitProviders(),
					Description: `Connection providers`,
				},
				"connection_id": &graphql.Field{
					Type:        graphql.String,
					Description: `unique connection_id provided by external OAuth`,
				},
				"accessToken": &graphql.Field{
					Type:        graphql.String,
					Description: `access token for external authorization`,
				},
				"refreshToken": &graphql.Field{
					Type:        graphql.String,
					Description: `for re-fetching access token`,
				},
				"accessTokenExpiry": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
				"refreshTokenExpiry": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
				"revoked": &graphql.Field{
					Type: graphql.Boolean,
				},
				"account_id": &graphql.Field{
					Type:        graphql.String,
					Description: `every connection must be associated with account_id`,
				},
				"user_name": &graphql.Field{
					Type:        graphql.String,
					Description: `specific user-name required while fetching the repo`,
				},
			},
		})
	}
	return gql__type_VCSConnection
}

func Gql__type_ListVCSConnectionResponse() *graphql.Object {
	if gql__type_ListVCSConnectionResponse == nil {
		gql__type_ListVCSConnectionResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_ListVCSConnectionResponse",
			Fields: graphql.Fields{
				"vcs_connection": &graphql.Field{
					Type: graphql.NewList(Gql__type_AccountVCSConnection()),
				},
			},
		})
	}
	return gql__type_ListVCSConnectionResponse
}

func Gql__type_ListVCSConnectionRequest() *graphql.Object {
	if gql__type_ListVCSConnectionRequest == nil {
		gql__type_ListVCSConnectionRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_ListVCSConnectionRequest",
			Fields: graphql.Fields{
				"account_id": &graphql.Field{
					Type: graphql.String,
				},
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
			},
		})
	}
	return gql__type_ListVCSConnectionRequest
}

func Gql__type_ListAllSupportedVCSProvidersResponse() *graphql.Object {
	if gql__type_ListAllSupportedVCSProvidersResponse == nil {
		gql__type_ListAllSupportedVCSProvidersResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_ListAllSupportedVCSProvidersResponse",
			Fields: graphql.Fields{
				"providers": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__type_ListAllSupportedVCSProvidersResponse
}

func Gql__type_CallbackRequest() *graphql.Object {
	if gql__type_CallbackRequest == nil {
		gql__type_CallbackRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_CallbackRequest",
			Fields: graphql.Fields{
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
				"state": &graphql.Field{
					Type: graphql.String,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"account_id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_CallbackRequest
}

func Gql__type_AuthorizeResponse() *graphql.Object {
	if gql__type_AuthorizeResponse == nil {
		gql__type_AuthorizeResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_AuthorizeResponse",
			Fields: graphql.Fields{
				"redirect_url": &graphql.Field{
					Type: graphql.String,
				},
				"temp_jwt_token": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_AuthorizeResponse
}

func Gql__type_AuthorizeRequest() *graphql.Object {
	if gql__type_AuthorizeRequest == nil {
		gql__type_AuthorizeRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_AuthorizeRequest",
			Fields: graphql.Fields{
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
				"label": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_AuthorizeRequest
}

func Gql__type_AccountVCSConnection() *graphql.Object {
	if gql__type_AccountVCSConnection == nil {
		gql__type_AccountVCSConnection = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_AccountVCSConnection",
			Fields: graphql.Fields{
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
				"account_id": &graphql.Field{
					Type: graphql.String,
				},
				"label": &graphql.Field{
					Type: graphql.String,
				},
				"user_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_AccountVCSConnection
}

func Gql__input_VCSConnection() *graphql.InputObject {
	if gql__input_VCSConnection == nil {
		gql__input_VCSConnection = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_VCSConnection",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Description: `id generated by uuid`,
					Type:        graphql.String,
				},
				"label": &graphql.InputObjectFieldConfig{
					Description: `label of VCS Connection`,
					Type:        graphql.String,
				},
				"provider": &graphql.InputObjectFieldConfig{
					Description: `Connection providers`,
					Type:        types.Gql__enum_GitProviders(),
				},
				"connection_id": &graphql.InputObjectFieldConfig{
					Description: `unique connection_id provided by external OAuth`,
					Type:        graphql.String,
				},
				"accessToken": &graphql.InputObjectFieldConfig{
					Description: `access token for external authorization`,
					Type:        graphql.String,
				},
				"refreshToken": &graphql.InputObjectFieldConfig{
					Description: `for re-fetching access token`,
					Type:        graphql.String,
				},
				"accessTokenExpiry": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
				"refreshTokenExpiry": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
				"revoked": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"account_id": &graphql.InputObjectFieldConfig{
					Description: `every connection must be associated with account_id`,
					Type:        graphql.String,
				},
				"user_name": &graphql.InputObjectFieldConfig{
					Description: `specific user-name required while fetching the repo`,
					Type:        graphql.String,
				},
			},
		})
	}
	return gql__input_VCSConnection
}

func Gql__input_ListVCSConnectionResponse() *graphql.InputObject {
	if gql__input_ListVCSConnectionResponse == nil {
		gql__input_ListVCSConnectionResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_ListVCSConnectionResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"vcs_connection": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_AccountVCSConnection()),
				},
			},
		})
	}
	return gql__input_ListVCSConnectionResponse
}

func Gql__input_ListVCSConnectionRequest() *graphql.InputObject {
	if gql__input_ListVCSConnectionRequest == nil {
		gql__input_ListVCSConnectionRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_ListVCSConnectionRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"account_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
			},
		})
	}
	return gql__input_ListVCSConnectionRequest
}

func Gql__input_ListAllSupportedVCSProvidersResponse() *graphql.InputObject {
	if gql__input_ListAllSupportedVCSProvidersResponse == nil {
		gql__input_ListAllSupportedVCSProvidersResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_ListAllSupportedVCSProvidersResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"providers": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__input_ListAllSupportedVCSProvidersResponse
}

func Gql__input_CallbackRequest() *graphql.InputObject {
	if gql__input_CallbackRequest == nil {
		gql__input_CallbackRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_CallbackRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"state": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"code": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"account_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_CallbackRequest
}

func Gql__input_AuthorizeResponse() *graphql.InputObject {
	if gql__input_AuthorizeResponse == nil {
		gql__input_AuthorizeResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_AuthorizeResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"redirect_url": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"temp_jwt_token": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_AuthorizeResponse
}

func Gql__input_AuthorizeRequest() *graphql.InputObject {
	if gql__input_AuthorizeRequest == nil {
		gql__input_AuthorizeRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_AuthorizeRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"label": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_AuthorizeRequest
}

func Gql__input_AccountVCSConnection() *graphql.InputObject {
	if gql__input_AccountVCSConnection == nil {
		gql__input_AccountVCSConnection = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_AccountVCSConnection",
			Fields: graphql.InputObjectConfigFieldMap{
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"account_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"label": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"user_name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_AccountVCSConnection
}

// graphql__resolver_VCSConnectionService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_VCSConnectionService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_VCSConnectionService creates pointer of service struct
func new_graphql_resolver_VCSConnectionService(conn *grpc.ClientConn) *graphql__resolver_VCSConnectionService {
	return &graphql__resolver_VCSConnectionService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_VCSConnectionService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_VCSConnectionService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"VCSConnections": &graphql.Field{
			Type: Gql__type_ListVCSConnectionResponse(),
			Args: graphql.FieldConfigArgument{
				"account_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"provider": &graphql.ArgumentConfig{
					Type: types.Gql__enum_GitProviders(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ListVCSConnectionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for VCSConnections")
				}
				client := NewVCSConnectionServiceClient(conn)
				resp, err := client.ListVCSConnection(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ListVCSConnection")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_VCSConnectionService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterVCSConnectionServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterVCSConnectionServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterVCSConnectionServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service VCSConnectionService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterVCSConnectionServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_VCSConnectionService(conn))
}
