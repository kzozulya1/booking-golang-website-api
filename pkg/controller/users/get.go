package users

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Get user data
var MeGet = func(w http.ResponseWriter, r *http.Request) {
	//Get User id from context
	uId, err := u.GetContextUserIdValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//Load user from DB
	user := &model.User{}
	err = user.Load(uId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//Get user by id
var GetOne = func(w http.ResponseWriter, r *http.Request) {
	//Get id param and convert into int
	params := mux.Vars(r)
	uId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "id param is invalid", http.StatusBadRequest)
		return
	}
	//Load user model
	user := &model.User{}
	err = user.Load(uint(uId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}