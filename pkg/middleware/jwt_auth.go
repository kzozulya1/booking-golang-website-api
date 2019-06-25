package middleware

import (
	"net/http"
	"context"
	"app/pkg/model"
	u "app/pkg/utils"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("api-token") //Grab the token from the header
		if tokenHeader != "" {
			//Token is set
			tk := &model.Token{}
			//Decode it, extract user id
			userId, err := tk.DecodeToken(tokenHeader)
			if err != nil {
				response = u.Message(false, err.Error())
				w.WriteHeader(http.StatusForbidden)
				u.RespondJson(w, response)
				return
			}

			//Save user it into HTTP request context
			ctx := context.WithValue(r.Context(), "user", userId)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) //proceed in the middleware chain
	})
}
