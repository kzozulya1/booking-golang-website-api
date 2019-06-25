package utils

import (
	"encoding/json"
	"net/http"
)

//Respond with error
func RespondError(w http.ResponseWriter, errorMsg string, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"error": errorMsg}
	json.NewEncoder(w).Encode(resp)
}
