// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"github.com/graphql-go/graphql"
	gql_ptypes_timestamp "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamp"
)

var (
	gql__enum_SOURCE                      *graphql.Enum        // enum SOURCE in user_profile_internal.proto
	gql__type_UserProfile                 *graphql.Object      // message UserProfile in user_profile_internal.proto
	gql__type_UpdateUserProfileRequest    *graphql.Object      // message UpdateUserProfileRequest in user_profile_internal.proto
	gql__type_GetUserProfileBySubRequest  *graphql.Object      // message GetUserProfileBySubRequest in user_profile_internal.proto
	gql__type_DeleteUserProfileRequest    *graphql.Object      // message DeleteUserProfileRequest in user_profile_internal.proto
	gql__type_CreateUserProfileRequest    *graphql.Object      // message CreateUserProfileRequest in user_profile_internal.proto
	gql__input_UserProfile                *graphql.InputObject // message UserProfile in user_profile_internal.proto
	gql__input_UpdateUserProfileRequest   *graphql.InputObject // message UpdateUserProfileRequest in user_profile_internal.proto
	gql__input_GetUserProfileBySubRequest *graphql.InputObject // message GetUserProfileBySubRequest in user_profile_internal.proto
	gql__input_DeleteUserProfileRequest   *graphql.InputObject // message DeleteUserProfileRequest in user_profile_internal.proto
	gql__input_CreateUserProfileRequest   *graphql.InputObject // message CreateUserProfileRequest in user_profile_internal.proto
)

func Gql__enum_SOURCE() *graphql.Enum {
	if gql__enum_SOURCE == nil {
		gql__enum_SOURCE = graphql.NewEnum(graphql.EnumConfig{
			Name: "Pb_Enum_SOURCE",
			Values: graphql.EnumValueConfigMap{
				"UNKNOWN": &graphql.EnumValueConfig{
					Value: SOURCE(0),
				},
				"GITHUB": &graphql.EnumValueConfig{
					Value: SOURCE(1),
				},
				"GITLAB": &graphql.EnumValueConfig{
					Value: SOURCE(2),
				},
				"BITBUCKET": &graphql.EnumValueConfig{
					Value: SOURCE(3),
				},
				"EMAIL_PASSWORD": &graphql.EnumValueConfig{
					Value: SOURCE(4),
				},
			},
		})
	}
	return gql__enum_SOURCE
}

func Gql__type_UserProfile() *graphql.Object {
	if gql__type_UserProfile == nil {
		gql__type_UserProfile = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UserProfile",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.String,
					Description: `generated by uuid`,
				},
				"sub": &graphql.Field{
					Type:        graphql.String,
					Description: `external unique_id`,
				},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: `name of the user`,
				},
				"user_name": &graphql.Field{
					Type:        graphql.String,
					Description: `a username provided by external application`,
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
					Type: Gql__enum_SOURCE(),
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

func Gql__type_UpdateUserProfileRequest() *graphql.Object {
	if gql__type_UpdateUserProfileRequest == nil {
		gql__type_UpdateUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UpdateUserProfileRequest",
			Fields: graphql.Fields{
				"user_profile": &graphql.Field{
					Type: Gql__type_UserProfile(),
				},
			},
		})
	}
	return gql__type_UpdateUserProfileRequest
}

func Gql__type_GetUserProfileBySubRequest() *graphql.Object {
	if gql__type_GetUserProfileBySubRequest == nil {
		gql__type_GetUserProfileBySubRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserProfileBySubRequest",
			Fields: graphql.Fields{
				"sub": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetUserProfileBySubRequest
}

func Gql__type_DeleteUserProfileRequest() *graphql.Object {
	if gql__type_DeleteUserProfileRequest == nil {
		gql__type_DeleteUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_DeleteUserProfileRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_DeleteUserProfileRequest
}

func Gql__type_CreateUserProfileRequest() *graphql.Object {
	if gql__type_CreateUserProfileRequest == nil {
		gql__type_CreateUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_CreateUserProfileRequest",
			Fields: graphql.Fields{
				"user_profile": &graphql.Field{
					Type: Gql__type_UserProfile(),
				},
			},
		})
	}
	return gql__type_CreateUserProfileRequest
}

func Gql__input_UserProfile() *graphql.InputObject {
	if gql__input_UserProfile == nil {
		gql__input_UserProfile = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UserProfile",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Description: `generated by uuid`,
					Type:        graphql.String,
				},
				"sub": &graphql.InputObjectFieldConfig{
					Description: `external unique_id`,
					Type:        graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Description: `name of the user`,
					Type:        graphql.String,
				},
				"user_name": &graphql.InputObjectFieldConfig{
					Description: `a username provided by external application`,
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
					Type: Gql__enum_SOURCE(),
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

func Gql__input_UpdateUserProfileRequest() *graphql.InputObject {
	if gql__input_UpdateUserProfileRequest == nil {
		gql__input_UpdateUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UpdateUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_profile": &graphql.InputObjectFieldConfig{
					Type: Gql__input_UserProfile(),
				},
			},
		})
	}
	return gql__input_UpdateUserProfileRequest
}

func Gql__input_GetUserProfileBySubRequest() *graphql.InputObject {
	if gql__input_GetUserProfileBySubRequest == nil {
		gql__input_GetUserProfileBySubRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserProfileBySubRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"sub": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetUserProfileBySubRequest
}

func Gql__input_DeleteUserProfileRequest() *graphql.InputObject {
	if gql__input_DeleteUserProfileRequest == nil {
		gql__input_DeleteUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_DeleteUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_DeleteUserProfileRequest
}

func Gql__input_CreateUserProfileRequest() *graphql.InputObject {
	if gql__input_CreateUserProfileRequest == nil {
		gql__input_CreateUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_CreateUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_profile": &graphql.InputObjectFieldConfig{
					Type: Gql__input_UserProfile(),
				},
			},
		})
	}
	return gql__input_CreateUserProfileRequest
}
