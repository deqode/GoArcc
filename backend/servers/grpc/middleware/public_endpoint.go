package middleware

// TODO: Priority zero, validate if these are available are registered, otherwise fail on startup
var publicEndpoint = []string{
	"/alfred.vcs_connection.v1.VCSConnections/ListAllSupportedVCSProviders",
	"/alfred.auth.v1.Authentications/Login",
	"/alfred.auth.v1.Authentications/LoginCallback",
	"/alfred.user_profile.v1.user-profile/GetUserProfileBySub",
}
