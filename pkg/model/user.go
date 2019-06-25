package model

import (
	conn "app/pkg/connection"
	"app/pkg/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

//const MIN_PASSWORD_LENGTH = 6

//User representation struct
type User struct {
	ID           uint          `gorm:"primary_key:true" json:"id"`
	Name         string        `json:"name"`
	Email        string        `gorm:"unique;not null" json:"email"`
	Password     string        `json:"password"`
	CreatedAt    string        `json:"created_at"`
	ModifiedAt   string        `json:"modified_at"`
	IsRoomMaster uint          `json:"is_room_master"`
	Age          uint          `json:"age"`
	Reservations []Reservation `gorm:"foreignkey:UserId" json:"reservations"`
	Rooms        []Room        `gorm:"foreignkey:OwnerId" json:"rooms"`
}

// Set DB table name
func (User) TableName() string {
	return "user"
}

//Check email and password are ok
func (u *User) validateEmailPwd() error {
	//Check email is ok, by RegExp
	emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailPattern.Match([]byte(u.Email)) {
		return fmt.Errorf("Email address is invalid")
	}
	//Optionally we can resrtict min pwd length

	//Check length is ok
	// if len(u.Password) < MIN_PASSWORD_LENGTH {
	// 	return fmt.Errorf("Min password length is %d", MIN_PASSWORD_LENGTH)
	// }
	return nil
}

//Generate encrypted password
func (u *User) genPassword(pwd string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashedPassword)
}

//Validate user data is ok, and user doesn't exist in DB yet
func (u *User) validateBeforeCreate() error {
	if err := u.validateEmailPwd(); err != nil {
		return err
	}
	//Check email is unique and doesn't exist
	tmpUser := &User{}
	//Check users doesn't exist
	if err := conn.GetDB().Where("email = ?", u.Email).First(tmpUser).Error; err != nil && err == gorm.ErrRecordNotFound {
		return nil
	} else {
		return fmt.Errorf("User with specified email already exists")
	}
}

// Create facility
// Return JWT token if success
func (u *User) Create() (string, error) {
	if err := u.validateBeforeCreate(); err != nil {
		return "", err
	}
	//Generate password
	u.Password = u.genPassword(u.Password)

	//Stamp with current datetime
	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.ModifiedAt = time.Now().Format("2006-01-02 15:04:05")

	//Create user in database
	if err := conn.GetDB().Create(u).Error; err != nil {
		return "", err
	}
	if u.ID <= 0 {
		return "", fmt.Errorf("Failed to create account")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: u.ID}
	tkString, err := tk.NewToken()
	if err != nil {
		return "", err
	}
	return tkString, nil
}

// Login facility
// Return JWT token if success
func (websiteUser *User) Login() (string, error) {
	//Check email and password are ok
	if err := websiteUser.validateEmailPwd(); err != nil {
		return "", err
	}

	dbUser := &User{}
	if err := conn.GetDB().Where("email = ?", websiteUser.Email).First(dbUser).Error; err != nil && err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("User with specified email not found")
	}

	//Check password
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(websiteUser.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return "", fmt.Errorf("Password is incorrect")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: dbUser.ID}
	tkString, err := tk.NewToken()
	if err != nil {
		return "", err
	}
	return tkString, nil
}

// Load user by ID
// Preload his reservsations and room he owns, in desc order
func (u *User) Load(userId uint) error {
	if err := conn.GetDB()/*.Debug()*/.Preload("Reservations", func(db *gorm.DB) *gorm.DB {
		return conn.GetDB().Order("id DESC")
	}).Preload("Rooms", func(db *gorm.DB) *gorm.DB {
		return conn.GetDB().Order("id DESC")
	}).Where("id = ?", userId).First(u).Error; err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("User not found")
	}
	//Security facility
	u.Password = ""

	//Date format for user's reservations
	if len(u.Reservations) > 0 {
		for key, _ := range u.Reservations {
			u.Reservations[key].CreatedAt = utils.TimeFormat(u.Reservations[key].CreatedAt)
			u.Reservations[key].BookFrom = utils.TimeFormat(u.Reservations[key].BookFrom)
			u.Reservations[key].BookTo = utils.TimeFormat(u.Reservations[key].BookTo)
		}
	}

	
	return nil
}

//Update user password
func (u *User) UpdatePassword(pwd string) {
	u.Password = u.genPassword(pwd)
}

//Stamp the datetime and persist user data
func (u *User) Save() error {
	u.ModifiedAt = time.Now().Format("2006-01-02 15:04:05")
	return conn.GetDB().Save(u).Error
}