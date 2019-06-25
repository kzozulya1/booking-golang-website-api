package model

import (
	"time"
	"fmt"
	conn "app/pkg/connection"
	"github.com/jinzhu/gorm"
	u "app/pkg/utils"
	"net/http"
)

//Struct for apartment representation
type Reservation struct {
	ID         uint      `gorm:"primary_key:true" json:"id"`
	RoomId     uint      `json:"room_id"`
	UserId     uint      `json:"user_id"`
	BookFrom   string    `json:"book_from"`
	BookTo     string    `json:"book_to"`
	CreatedAt  string    `json:"created_at"`
	ModifiedAt string    `gorm:"default:null" json:"modified_at"`
	Notes      string    `json:"notes"`
}

// Set DB table name
func (Reservation) TableName() string {
	return "reservation"
}

// Load apartment reservation by ID, sort by id desc
func (r *Reservation) Load(reservId uint) error {
	if err := conn.GetDB().Where("id = ?", reservId).First(r).Error; err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Apartment reservation not found")
	}
	return nil
}

//Delete apartment reservation
func (r *Reservation) Delete() error {
	return conn.GetDB().Delete(r).Error
}

//Validate apartment reservation data is ok
//Check apartment have not already reserved within requested dates
func (r *Reservation) validateBeforeCreate() (error, int) {
	//Check from/to dates are not empty
	if r.BookFrom == "" {
		return fmt.Errorf("Book from date is empty"), http.StatusBadRequest
	}
	if r.BookTo == "" {
		return fmt.Errorf("Book to date is empty"), http.StatusBadRequest
	}
	//Check already reserved
	tmpReservation := &Reservation{}
	if err := conn.GetDB()/*.Debug()*/.Where("(book_from <= ?) AND (book_to >= ?) AND room_id = ?", 
			r.BookTo, 
			r.BookFrom,
			r.RoomId,
		).First(tmpReservation).Error; err != nil && err == gorm.ErrRecordNotFound {
		  return nil, http.StatusOK
	}else{
		return fmt.Errorf("Already reserved from %s to %s",
			u.TimeFormat(tmpReservation.BookFrom),
			u.TimeFormat(tmpReservation.BookTo),
			), http.StatusConflict
	}
}

// Apartment reservation book facility
func (r *Reservation) Create() (error, int) {
	err, httpStatus := r.validateBeforeCreate();
	if  err != nil {
		return err, httpStatus
	}
	//Stamp with current datetime
	r.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	r.ModifiedAt = time.Now().Format("2006-01-02 15:04:05")

	//Create apartment reservation in database
	if err := conn.GetDB().Create(r).Error; err != nil {
		return err, http.StatusInternalServerError
	}
	if r.ID <= 0 {
		return fmt.Errorf("Failed to create apartment reservation"), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}