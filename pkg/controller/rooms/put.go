package rooms

import (
	"app/pkg/model"
	u "app/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Update apartment
var Update = func(w http.ResponseWriter, r *http.Request) {
	//Get user id from context value
	uId, err := u.GetContextUserIdValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//Get param and validate
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "id param is invalid", http.StatusBadRequest)
		return
	}
	//Get data from frontend
	websiteRoom := &model.Room{}
	json.NewDecoder(r.Body).Decode(websiteRoom)

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

	//Copy data
	copyData(websiteRoom, dbRoom)

	//Persist new data
	err = dbRoom.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	u.RespondEncodeStructInJson(w,dbRoom)
	
	// resp := u.Message(true, "Apartment successfully updated")
	// u.RespondJson(w, resp)
}

//Copy data
func copyData(srcR *model.Room, dstR *model.Room) {
	dstR.Description = srcR.Description
	dstR.BodyCapacity = srcR.BodyCapacity
	dstR.Address = srcR.Address
	dstR.AllowSmoking = srcR.AllowSmoking
	dstR.AllowParking = srcR.AllowParking
	dstR.AllowChildren = srcR.AllowChildren
	dstR.Img1 = srcR.Img1
	dstR.Img2 = srcR.Img2
	dstR.Img3 = srcR.Img3
	dstR.ImgMain = srcR.ImgMain
}
