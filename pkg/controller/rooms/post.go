package rooms

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	//"fmt"
	"net/http"
)

//Create an apartment
var Create = func(w http.ResponseWriter, r *http.Request) {
	//Get user id from context value
	uId, err := u.GetContextUserIdValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//Get posted user data
	room := &model.Room{}
	json.NewDecoder(r.Body).Decode(room)

	//Assign owner ID
	room.OwnerId = uId

	//Create apartment
	err = room.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	u.RespondEncodeStructInJson(w,room)
	
	// jsonData, err := json.Marshal(room)
	// if err != nil{
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(jsonData)
}