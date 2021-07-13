package middleware

// TODO: Priority zero, validate if these are available are registered, otherwise fail on startup
var publicEndpoint = []string{
	"/goarcc.vcs_connection.v1.VCSConnections/ListAllSupportedVCSProviders",
	"/goarcc.authentication.v1.Authentications/Login",
	"/goarcc.authentication.v1.Authentications/LoginCallback",
	"/goarcc.user_profile.v1.user-profile/GetUserProfileBySub",
}
