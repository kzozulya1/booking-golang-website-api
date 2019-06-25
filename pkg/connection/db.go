package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var db *gorm.DB

//Establ. postgres connection
func init() {
	dbUri := os.Getenv("DATABASE_CONN")
	fmt.Println("Connecting to ", dbUri, "...")

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}
	db = conn
}

//Get DB connection
func GetDB() *gorm.DB {
	return db
}
