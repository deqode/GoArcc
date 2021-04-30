package middleware

import (
	"google.golang.org/grpc/metadata"
	"net/http"
)

func ChangeContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request = request.WithContext(metadata.AppendToOutgoingContext(request.Context(), "Authorization", request.Header.Get("Authorization")))
		h.ServeHTTP(writer, request)
	})
}
