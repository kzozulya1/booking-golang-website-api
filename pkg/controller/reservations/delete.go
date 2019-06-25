package reservations

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Delete the araptment
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
	//Load apartment reservation model
	reservation := &model.Reservation{}
	err = reservation.Load(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Check current user owns the apartment
	if reservation.UserId != uId {
		http.Error(w, fmt.Sprintf("You don't own the apartment reservation with ID = %d", id), http.StatusInternalServerError)
		return
	}
	//Persist new data
	err = reservation.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
