package rooms

import (
	"app/pkg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Get all apartments, no auth token required
var GetAll = func(w http.ResponseWriter, r *http.Request) {
	//Get apartment collection
	rooms, err := model.GetAllRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

//Get apartment by id, no auth token required
var GetOne = func(w http.ResponseWriter, r *http.Request) {
	//Get id param and validate
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "id param is invalid", http.StatusBadRequest)
		return
	}
	//Find record
	room := model.Room{}
	err = room.Load(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}