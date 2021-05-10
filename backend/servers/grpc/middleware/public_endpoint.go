package middleware

// TODO: Priority zero, validate if these are available are registered, otherwise fail on startup
var publicEndpoint = []string{
	"/alfred.vcs_connection.v1.vcs-connection/ListAllSupportedVCSProviders",
	"/alfred.authentication.v1.UserLoginService/UserLogin",
	"/alfred.authentication.v1.UserLoginService/UserLoginCallback",
	"/alfred.user_profile.v1.user-profile/GetUserProfileBySub",
	//"/alfred.user_profile.v1.user-profile/CreateUserProfile",
}
