package model

import (
	conn "app/pkg/connection"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//Apartment representation
type Room struct {
	ID                             uint          `gorm:"primary_key:true" json:"id"`
	Description                    string        `json:"description"`
	BodyCapacity                   uint          `json:"body_capacity"`
	OwnerId                        uint          `json:"owner_id"`
	AllowSmoking                   bool          `json:"allow_smoking"`
	AllowParking                   bool          `json:"allow_parking"`
	AllowChildren                  bool          `json:"allow_children"`
	Address                        string        `json:"address"`
	CreatedAt                      string        `json:"created_at"`
	ModifiedAt                     string        `gorm:"default:null" json:"modified_at"`
	Img1                           string        `gorm:"column:img_1" json:"img_1"`
	Img2                           string        `gorm:"column:img_2" json:"img_2"`
	Img3                           string        `gorm:"column:img_3" json:"img_3"`
	ImgMain                        string        `json:"img_main"`
	Reservations                   []Reservation `gorm:"foreignkey:RoomId" json:"reservations"`
}

// Set DB table name
func (Room) TableName() string {
	return "room"
}

// Load apartment by ID, sort by id desc
func (r *Room) Load(roomId uint) error {
	if err := conn.GetDB()/*.Debug()*/.Preload("Reservations", func(db *gorm.DB) *gorm.DB {
		return conn.GetDB().Order("id DESC")
	}).Where("id = ?", roomId).First(r).Error; err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Apartment not found")
	}
	return nil
}

//Validate apartment data is ok
func (r *Room) validateBeforeCreate() error {
	if r.Description == "" {
		return fmt.Errorf("Apartment description is empty")
	}
	if r.BodyCapacity == 0 {
		return fmt.Errorf("Apartment guest capacity is empty")
	}
	return nil
}

// Apartment create facility
func (r *Room) Create() error {
	if err := r.validateBeforeCreate(); err != nil {
		return err
	}
	//Stamp with current datetime
	r.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	r.ModifiedAt = time.Now().Format("2006-01-02 15:04:05")

	//Create apartment in database
	if err := conn.GetDB().Create(r).Error; err != nil {
		return err
	}
	if r.ID <= 0 {
		return fmt.Errorf("Failed to create apartment")
	}
	return nil
}

//Persist apartment data
func (r *Room) Save() error {
	r.ModifiedAt = time.Now().Format("2006-01-02 15:04:05")
	return conn.GetDB().Save(r).Error
}

//Delete apartment
func (r *Room) Delete() error {
	return conn.GetDB().Delete(r).Error
}

// COLLECTION
//Get all apartments, ordered by id desc
//old order is modified_at desc, id
func GetAllRooms() ([]Room, error) {
	var rooms = make([]Room, 1)
	err := conn.GetDB().Preload("Reservations", func(db *gorm.DB) *gorm.DB {
		return conn.GetDB().Order("id DESC")
	}).Order("id desc").Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}