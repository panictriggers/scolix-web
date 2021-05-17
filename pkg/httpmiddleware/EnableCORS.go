package httpmiddleware

import "net/http"

// EnableCors sets the Access Control Allow Origin header to '*' to allow requests from a different domain
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		},
	)
}