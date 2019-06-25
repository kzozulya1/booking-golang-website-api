package users

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	"net/http"
)

//Sing up user
var Create = func(w http.ResponseWriter, r *http.Request) {
	//Get posted user data
	user := &model.User{}
	json.NewDecoder(r.Body).Decode(user)

	//Create user
	jwtToken, err := user.Create()
	if err != nil {
		u.RespondError(w,err.Error(),http.StatusBadRequest)
		return
	}
	resp := u.Message(true, "User successfully created")
	resp["jwt_token"] = jwtToken
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}


//Sign in user
var Login = func(w http.ResponseWriter, r *http.Request) {
	//Get posted user data
	user := &model.User{}
	json.NewDecoder(r.Body).Decode(user)
	//Login user
	jwtToken, err := user.Login()
	if err != nil {
		//!!!
		//http.Error(w, err.Error(), http.StatusBadRequest)
		u.RespondError(w,err.Error(),http.StatusBadRequest)
		//!!!
		return
	}
	resp := u.Message(true, "User successfully logged in")
	resp["jwt_token"] = jwtToken
	u.RespondJson(w, resp)
}