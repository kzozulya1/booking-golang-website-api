package rooms

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Delete the apartment
var Delete = func(w http.ResponseWriter, r *http.Request) {
	//Get user id from context value
	uId, err := u.GetContextUserIdValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//Get id param and validate
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "id param is invalid", http.StatusBadRequest)
		return
	}
	//Load DB apartment model
	dbRoom := &model.Room{}
	err = dbRoom.Load(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Check current user owns an apartment
	if dbRoom.OwnerId != uId {
		http.Error(w, fmt.Sprintf("You don't own the apartment with ID = %d", id), http.StatusInternalServerError)
		return
	}
	//Persist new data
	err = dbRoom.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
