package httpmiddleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

// Log the HTTP requests that go through this router
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lrw := negroni.NewResponseWriter(w)

			next.ServeHTTP(lrw, r)
			fmt.Printf("[" + time.Now().Format(time.UnixDate) + "] " + r.RemoteAddr + " [" + r.Method + "] %d " + r.Host + r.URL.String() + "\n", lrw.Status())
		},
	)
}
