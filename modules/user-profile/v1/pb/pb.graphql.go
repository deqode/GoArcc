// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	types "goarcc/protos/types"
	"google.golang.org/grpc"

	gql_ptypes_timestamp "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamp"

	gql_ptypes_empty "github.com/ysugimoto/grpc-graphql-gateway/ptypes/empty"
)

var (
	gql__type_UserProfile                 *graphql.Object      // message UserProfile in pb/user_profile.proto
	gql__type_GetUserProfileBySubRequest  *graphql.Object      // message GetUserProfileBySubRequest in pb/user_profile.proto
	gql__input_UserProfile                *graphql.InputObject // message UserProfile in pb/user_profile.proto
	gql__input_GetUserProfileBySubRequest *graphql.InputObject // message GetUserProfileBySubRequest in pb/user_profile.proto
)

func Gql__type_UserProfile() *graphql.Object {
	if gql__type_UserProfile == nil {
		gql__type_UserProfile = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UserProfile",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"sub": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `external unique_id provided by vcs provider`,
				},
				"name": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `name of the user`,
				},
				"user_name": &graphql.Field{
					Type:        graphql.String,
					Description: `a username provided by vcs provider`,
				},
				"email": &graphql.Field{
					Type:        graphql.String,
					Description: `email of user`,
				},
				"phone_number": &graphql.Field{
					Type:        graphql.String,
					Description: `phone of user`,
				},
				"external_source": &graphql.Field{
					Type: types.Gql__enum_VCSProviders(),
				},
				"profile_pic_url": &graphql.Field{
					Type: graphql.String,
				},
				"token_valid_till": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
			},
		})
	}
	return gql__type_UserProfile
}

func Gql__type_GetUserProfileBySubRequest() *graphql.Object {
	if gql__type_GetUserProfileBySubRequest == nil {
		gql__type_GetUserProfileBySubRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserProfileBySubRequest",
			Fields: graphql.Fields{
				"sub": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_GetUserProfileBySubRequest
}

func Gql__input_UserProfile() *graphql.InputObject {
	if gql__input_UserProfile == nil {
		gql__input_UserProfile = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UserProfile",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"sub": &graphql.InputObjectFieldConfig{
					Description: `external unique_id provided by vcs provider`,
					Type:        graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.InputObjectFieldConfig{
					Description: `name of the user`,
					Type:        graphql.NewNonNull(graphql.String),
				},
				"user_name": &graphql.InputObjectFieldConfig{
					Description: `a username provided by vcs provider`,
					Type:        graphql.String,
				},
				"email": &graphql.InputObjectFieldConfig{
					Description: `email of user`,
					Type:        graphql.String,
				},
				"phone_number": &graphql.InputObjectFieldConfig{
					Description: `phone of user`,
					Type:        graphql.String,
				},
				"external_source": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_VCSProviders(),
				},
				"profile_pic_url": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"token_valid_till": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
			},
		})
	}
	return gql__input_UserProfile
}

func Gql__input_GetUserProfileBySubRequest() *graphql.InputObject {
	if gql__input_GetUserProfileBySubRequest == nil {
		gql__input_GetUserProfileBySubRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserProfileBySubRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"sub": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_GetUserProfileBySubRequest
}

// graphql__resolver_UserProfiles is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_UserProfiles struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_UserProfiles creates pointer of service struct
func new_graphql_resolver_UserProfiles(conn *grpc.ClientConn) *graphql__resolver_UserProfiles {
	return &graphql__resolver_UserProfiles{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_UserProfiles) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_UserProfiles) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"user": &graphql.Field{
			Type: Gql__type_UserProfile(),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req gql_ptypes_empty.Empty
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for user")
				}
				client := NewUserProfilesClient(conn)
				resp, err := client.GetUserProfile(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUserProfile")
				}
				return resp, nil
			},
		},
		"userBySub": &graphql.Field{
			Type: Gql__type_UserProfile(),
			Args: graphql.FieldConfigArgument{
				"sub": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetUserProfileBySubRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for userBySub")
				}
				client := NewUserProfilesClient(conn)
				resp, err := client.GetUserProfileBySub(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUserProfileBySub")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_UserProfiles) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterUserProfilesGraphqlHandler with *grpc.ClientConn manually.
func RegisterUserProfilesGraphql(mux *runtime.ServeMux) error {
	return RegisterUserProfilesGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service UserProfiles {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterUserProfilesGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_UserProfiles(conn))
}