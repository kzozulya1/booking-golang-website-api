package middleware

import (
	"net/http"
)

//Apply CORS politics
var Cors = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, api-token")
		h.Set("Access-Control-Allow-Methods", "HEAD, GET, PUT, POST, OPTIONS, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r) //proceed in the middleware chain
	})
}