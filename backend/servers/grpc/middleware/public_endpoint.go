package middleware

// TODO: Priority zero, validate if these are available are registered, otherwise fail on startup
var publicEndpoint = []string{
	"/alfred.vcs_connection.v1.VCSConnection/ListAllSupportedVCSProviders",
	"/alfred.auth.v1.UserLoginService/UserLogin",
	"/alfred.auth.v1.UserLoginService/UserLoginCallback",
	"/alfred.user_profile.v1.UserProfileService/GetUserProfileBySub",
	//"/alfred.user_profile.v1.UserProfileService/CreateUserProfile",
}
