package utils

import (
	"fmt"
	"net/http"
)

//Get context user value. Setup in pkg/middleware/jwt_auth.go
func GetContextUserIdValue(r *http.Request) (uint, error) {
	userContext := r.Context().Value("user")
	if userContext == nil {
		return 0, fmt.Errorf("HTTP context value 'user' is not defined")
	}
	return userContext.(uint), nil
}