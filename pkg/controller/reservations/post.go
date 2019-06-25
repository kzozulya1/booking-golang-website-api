package reservations

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	
	"strconv"
	//"fmt"
	"github.com/gorilla/mux"	
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
	//Get id param and convert into int
	params := mux.Vars(r)
	roomId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "id param is invalid", http.StatusBadRequest)
		return
	}

	//Get posted user data
	reservation := &model.Reservation{}
	//Store book_from, book_to, notes into  Reservation struct
	json.NewDecoder(r.Body).Decode(reservation)

	//Assign apartment guest id and apartment id to struct fields
	reservation.RoomId = uint(roomId)
	reservation.UserId = uId

	//Create apartment
	err, httpStatus := reservation.Create()
	if err != nil {
		// w.WriteHeader(httpStatus)
		// w.Header().Set("Content-Type", "application/json")
		// resp := map[string]string{"error": err.Error()}
		// json.NewEncoder(w).Encode(resp)
		u.RespondError(w,err.Error(),httpStatus)
		return
	}

	resp := u.Message(true, "Apartment reservation successfully created")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
