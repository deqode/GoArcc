// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_MachineConfig                    *graphql.Object      // message MachineConfig in AwsArchitectureService.proto
	gql__type_GetAwsMachineTypesResponse       *graphql.Object      // message GetAwsMachineTypesResponse in AwsArchitectureService.proto
	gql__type_GetAwsMachineTypesRequest        *graphql.Object      // message GetAwsMachineTypesRequest in AwsArchitectureService.proto
	gql__type_GetAwsLaunchTypesResponse        *graphql.Object      // message GetAwsLaunchTypesResponse in AwsArchitectureService.proto
	gql__type_GetAwsLaunchTypesRequest         *graphql.Object      // message GetAwsLaunchTypesRequest in AwsArchitectureService.proto
	gql__type_GetAwsInstanceTypesResponse      *graphql.Object      // message GetAwsInstanceTypesResponse in AwsArchitectureService.proto
	gql__type_GetAwsInstanceTypesRequest       *graphql.Object      // message GetAwsInstanceTypesRequest in AwsArchitectureService.proto
	gql__type_GetAwsContainerServicesResponse  *graphql.Object      // message GetAwsContainerServicesResponse in AwsArchitectureService.proto
	gql__type_GetAwsContainerServicesRequest   *graphql.Object      // message GetAwsContainerServicesRequest in AwsArchitectureService.proto
	gql__input_MachineConfig                   *graphql.InputObject // message MachineConfig in AwsArchitectureService.proto
	gql__input_GetAwsMachineTypesResponse      *graphql.InputObject // message GetAwsMachineTypesResponse in AwsArchitectureService.proto
	gql__input_GetAwsMachineTypesRequest       *graphql.InputObject // message GetAwsMachineTypesRequest in AwsArchitectureService.proto
	gql__input_GetAwsLaunchTypesResponse       *graphql.InputObject // message GetAwsLaunchTypesResponse in AwsArchitectureService.proto
	gql__input_GetAwsLaunchTypesRequest        *graphql.InputObject // message GetAwsLaunchTypesRequest in AwsArchitectureService.proto
	gql__input_GetAwsInstanceTypesResponse     *graphql.InputObject // message GetAwsInstanceTypesResponse in AwsArchitectureService.proto
	gql__input_GetAwsInstanceTypesRequest      *graphql.InputObject // message GetAwsInstanceTypesRequest in AwsArchitectureService.proto
	gql__input_GetAwsContainerServicesResponse *graphql.InputObject // message GetAwsContainerServicesResponse in AwsArchitectureService.proto
	gql__input_GetAwsContainerServicesRequest  *graphql.InputObject // message GetAwsContainerServicesRequest in AwsArchitectureService.proto
)

func Gql__type_MachineConfig() *graphql.Object {
	if gql__type_MachineConfig == nil {
		gql__type_MachineConfig = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_MachineConfig",
			Fields: graphql.Fields{
				"parent": &graphql.Field{
					Type: graphql.String,
				},
				"instance_size": &graphql.Field{
					Type: graphql.String,
				},
				"v_cpu": &graphql.Field{
					Type: graphql.String,
				},
				"memory_in_gib": &graphql.Field{
					Type: graphql.String,
				},
				"instance_storage": &graphql.Field{
					Type: graphql.String,
				},
				"instance_storage_gib": &graphql.Field{
					Type: graphql.String,
				},
				"network_bandwidth_gbpc": &graphql.Field{
					Type: graphql.String,
				},
				"ebs_bandwidth_mbps": &graphql.Field{
					Type: graphql.String,
				},
				"baseline_performance": &graphql.Field{
					Type: graphql.String,
				},
				"cpu_credits_earned_hr": &graphql.Field{
					Type: graphql.String,
				},
				"network_burst_bandwidth_gbps": &graphql.Field{
					Type: graphql.String,
				},
				"ebs_burst_bandwidth_mbps": &graphql.Field{
					Type: graphql.String,
				},
				"mem": &graphql.Field{
					Type: graphql.String,
				},
				"mem_gib": &graphql.Field{
					Type: graphql.String,
				},
				"network_performance": &graphql.Field{
					Type: graphql.String,
				},
				"storage": &graphql.Field{
					Type: graphql.String,
				},
				"storage_gb": &graphql.Field{
					Type: graphql.String,
				},
				"network_performance_gib": &graphql.Field{
					Type: graphql.String,
				},
				"network_bandwidth": &graphql.Field{
					Type: graphql.String,
				},
				"ebs_bandwidth": &graphql.Field{
					Type: graphql.String,
				},
				"gpu_mem_gib": &graphql.Field{
					Type: graphql.String,
				},
				"gpu_mem": &graphql.Field{
					Type: graphql.String,
				},
				"gpu_p2p": &graphql.Field{
					Type: graphql.String,
				},
				"dedicated_ebs_bandwidth": &graphql.Field{
					Type: graphql.String,
				},
				"local_storage": &graphql.Field{
					Type: graphql.String,
				},
				"local_storage_gb": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_MachineConfig
}

func Gql__type_GetAwsMachineTypesResponse() *graphql.Object {
	if gql__type_GetAwsMachineTypesResponse == nil {
		gql__type_GetAwsMachineTypesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsMachineTypesResponse",
			Fields: graphql.Fields{
				"machine_types": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
				"machine_config": &graphql.Field{
					Type: graphql.NewList(Gql__type_MachineConfig()),
				},
			},
		})
	}
	return gql__type_GetAwsMachineTypesResponse
}

func Gql__type_GetAwsMachineTypesRequest() *graphql.Object {
	if gql__type_GetAwsMachineTypesRequest == nil {
		gql__type_GetAwsMachineTypesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsMachineTypesRequest",
			Fields: graphql.Fields{
				"instance_type": &graphql.Field{
					Type: graphql.String,
				},
				"parent": &graphql.Field{
					Type:        graphql.String,
					Description: `global parent : like Ecs , Eks`,
				},
			},
		})
	}
	return gql__type_GetAwsMachineTypesRequest
}

func Gql__type_GetAwsLaunchTypesResponse() *graphql.Object {
	if gql__type_GetAwsLaunchTypesResponse == nil {
		gql__type_GetAwsLaunchTypesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsLaunchTypesResponse",
			Fields: graphql.Fields{
				"launch_type": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__type_GetAwsLaunchTypesResponse
}

func Gql__type_GetAwsLaunchTypesRequest() *graphql.Object {
	if gql__type_GetAwsLaunchTypesRequest == nil {
		gql__type_GetAwsLaunchTypesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsLaunchTypesRequest",
			Fields: graphql.Fields{
				"container_service": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetAwsLaunchTypesRequest
}

func Gql__type_GetAwsInstanceTypesResponse() *graphql.Object {
	if gql__type_GetAwsInstanceTypesResponse == nil {
		gql__type_GetAwsInstanceTypesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsInstanceTypesResponse",
			Fields: graphql.Fields{
				"instance_types": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__type_GetAwsInstanceTypesResponse
}

func Gql__type_GetAwsInstanceTypesRequest() *graphql.Object {
	if gql__type_GetAwsInstanceTypesRequest == nil {
		gql__type_GetAwsInstanceTypesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsInstanceTypesRequest",
			Fields: graphql.Fields{
				"parent": &graphql.Field{
					Type: graphql.String,
				},
				"launch_type": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetAwsInstanceTypesRequest
}

func Gql__type_GetAwsContainerServicesResponse() *graphql.Object {
	if gql__type_GetAwsContainerServicesResponse == nil {
		gql__type_GetAwsContainerServicesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsContainerServicesResponse",
			Fields: graphql.Fields{
				"amazon_container_services": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__type_GetAwsContainerServicesResponse
}

func Gql__type_GetAwsContainerServicesRequest() *graphql.Object {
	if gql__type_GetAwsContainerServicesRequest == nil {
		gql__type_GetAwsContainerServicesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetAwsContainerServicesRequest",
			Fields: graphql.Fields{
				"nothing": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetAwsContainerServicesRequest
}

func Gql__input_MachineConfig() *graphql.InputObject {
	if gql__input_MachineConfig == nil {
		gql__input_MachineConfig = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_MachineConfig",
			Fields: graphql.InputObjectConfigFieldMap{
				"parent": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"instance_size": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"v_cpu": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"memory_in_gib": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"instance_storage": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"instance_storage_gib": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"network_bandwidth_gbpc": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"ebs_bandwidth_mbps": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"baseline_performance": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"cpu_credits_earned_hr": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"network_burst_bandwidth_gbps": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"ebs_burst_bandwidth_mbps": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"mem": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"mem_gib": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"network_performance": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"storage": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"storage_gb": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"network_performance_gib": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"network_bandwidth": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"ebs_bandwidth": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"gpu_mem_gib": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"gpu_mem": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"gpu_p2p": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"dedicated_ebs_bandwidth": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"local_storage": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"local_storage_gb": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_MachineConfig
}

func Gql__input_GetAwsMachineTypesResponse() *graphql.InputObject {
	if gql__input_GetAwsMachineTypesResponse == nil {
		gql__input_GetAwsMachineTypesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsMachineTypesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"machine_types": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				"machine_config": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_MachineConfig()),
				},
			},
		})
	}
	return gql__input_GetAwsMachineTypesResponse
}

func Gql__input_GetAwsMachineTypesRequest() *graphql.InputObject {
	if gql__input_GetAwsMachineTypesRequest == nil {
		gql__input_GetAwsMachineTypesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsMachineTypesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"instance_type": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"parent": &graphql.InputObjectFieldConfig{
					Description: `global parent : like Ecs , Eks`,
					Type:        graphql.String,
				},
			},
		})
	}
	return gql__input_GetAwsMachineTypesRequest
}

func Gql__input_GetAwsLaunchTypesResponse() *graphql.InputObject {
	if gql__input_GetAwsLaunchTypesResponse == nil {
		gql__input_GetAwsLaunchTypesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsLaunchTypesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"launch_type": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__input_GetAwsLaunchTypesResponse
}

func Gql__input_GetAwsLaunchTypesRequest() *graphql.InputObject {
	if gql__input_GetAwsLaunchTypesRequest == nil {
		gql__input_GetAwsLaunchTypesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsLaunchTypesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"container_service": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetAwsLaunchTypesRequest
}

func Gql__input_GetAwsInstanceTypesResponse() *graphql.InputObject {
	if gql__input_GetAwsInstanceTypesResponse == nil {
		gql__input_GetAwsInstanceTypesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsInstanceTypesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"instance_types": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__input_GetAwsInstanceTypesResponse
}

func Gql__input_GetAwsInstanceTypesRequest() *graphql.InputObject {
	if gql__input_GetAwsInstanceTypesRequest == nil {
		gql__input_GetAwsInstanceTypesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsInstanceTypesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"parent": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"launch_type": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetAwsInstanceTypesRequest
}

func Gql__input_GetAwsContainerServicesResponse() *graphql.InputObject {
	if gql__input_GetAwsContainerServicesResponse == nil {
		gql__input_GetAwsContainerServicesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsContainerServicesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"amazon_container_services": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__input_GetAwsContainerServicesResponse
}

func Gql__input_GetAwsContainerServicesRequest() *graphql.InputObject {
	if gql__input_GetAwsContainerServicesRequest == nil {
		gql__input_GetAwsContainerServicesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetAwsContainerServicesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"nothing": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetAwsContainerServicesRequest
}

// graphql__resolver_AwsService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_AwsService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_AwsService creates pointer of service struct
func new_graphql_resolver_AwsService(conn *grpc.ClientConn) *graphql__resolver_AwsService {
	return &graphql__resolver_AwsService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_AwsService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_AwsService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"getAwsContainerServices": &graphql.Field{
			Type: Gql__type_GetAwsContainerServicesResponse(),
			Args: graphql.FieldConfigArgument{
				"nothing": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetAwsContainerServicesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getAwsContainerServices")
				}
				client := NewAwsServiceClient(conn)
				resp, err := client.GetAwsContainerServices(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetAwsContainerServices")
				}
				return resp, nil
			},
		},
		"getAwsLaunchTypes": &graphql.Field{
			Type: Gql__type_GetAwsLaunchTypesResponse(),
			Args: graphql.FieldConfigArgument{
				"container_service": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetAwsLaunchTypesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getAwsLaunchTypes")
				}
				client := NewAwsServiceClient(conn)
				resp, err := client.GetAwsLaunchTypes(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetAwsLaunchTypes")
				}
				return resp, nil
			},
		},
		"getAwsInstanceTypes": &graphql.Field{
			Type: Gql__type_GetAwsInstanceTypesResponse(),
			Args: graphql.FieldConfigArgument{
				"parent": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"launch_type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetAwsInstanceTypesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getAwsInstanceTypes")
				}
				client := NewAwsServiceClient(conn)
				resp, err := client.GetAwsInstanceTypes(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetAwsInstanceTypes")
				}
				return resp, nil
			},
		},
		"getAwsMachineTypes": &graphql.Field{
			Type: Gql__type_GetAwsMachineTypesResponse(),
			Args: graphql.FieldConfigArgument{
				"instance_type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"parent": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: `global parent : like Ecs , Eks`,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetAwsMachineTypesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getAwsMachineTypes")
				}
				client := NewAwsServiceClient(conn)
				resp, err := client.GetAwsMachineTypes(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetAwsMachineTypes")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_AwsService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterAwsServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterAwsServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterAwsServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service AwsService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterAwsServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_AwsService(conn))
}