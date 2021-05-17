package httpmiddleware

import "net/http"

// Set the Content-Type to application/json for the API routes
func SetResponseTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}