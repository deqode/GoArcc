package middleware

import (
	"google.golang.org/grpc/metadata"
	"net/http"
)

// ChangeContext : ChangeContext is responsible to add the authorization token to the context so that
// in grpc server we can easily get auth token and verify the auth token.
func ChangeContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request = request.WithContext(metadata.AppendToOutgoingContext(request.Context(), "Authorization", request.Header.Get("Authorization")))
		h.ServeHTTP(writer, request)
	})
}
