package middleware

import "net/http"

//setupCorsResponse : is the helper method which is responsible to set the header in the response on every request.
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// AddCors : AddCors is the middleware which will add the cors in the response and response
// with status ok on preflight request.
func AddCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		setupCorsResponse(&writer, request)
		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
		}
		h.ServeHTTP(writer, request)
	})
}
