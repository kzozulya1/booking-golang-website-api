package utils

import (
	"encoding/json"
	"net/http"
)

//Respond with error
func RespondEncodeStructInJson(w http.ResponseWriter, structedData interface{}) {
	// w.WriteHeader(httpStatus)
	// w.Header().Set("Content-Type", "application/json")
	// resp := map[string]string{"error": errorMsg}
	// json.NewEncoder(w).Encode(resp)

	jsonData, err := json.Marshal(structedData)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
