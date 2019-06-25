package utils

import (
	"encoding/json"
	"net/http"
)

//Respond with json
func RespondJson(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
