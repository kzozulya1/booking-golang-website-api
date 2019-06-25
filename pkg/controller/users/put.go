package users

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	"net/http"
)

//Update user data
var MePut = func(w http.ResponseWriter, r *http.Request) {
	//Get user id from context value
	uId, err := u.GetContextUserIdValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//Get data from frontend
	websiteUser := &model.User{}
	json.NewDecoder(r.Body).Decode(websiteUser)
	if websiteUser.Name == "" {
		http.Error(w, "No name value specified", http.StatusUnauthorized)
		return
	}
	//Load db user model
	dbUser := &model.User{}
	err = dbUser.Load(uId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Copy website user data to db user struct
	dbUser.Name = websiteUser.Name
	if websiteUser.Password != "" {
		dbUser.UpdatePassword(websiteUser.Password)
	}
	//Persist new data
	err = dbUser.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbUser.Password = ""
	u.RespondEncodeStructInJson(w,dbUser)
}