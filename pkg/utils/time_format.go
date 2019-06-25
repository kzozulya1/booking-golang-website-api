package utils

import (
	"fmt"
	"time"
)

//User friendly time formatting
func TimeFormat(pgTime string) string {
	t, err := time.Parse(time.RFC3339, pgTime)
	if err != nil {
		fmt.Println("Error in time format: ", err.Error())
		return pgTime
	}
	//return t.Format("2 Jan 2006 15:04:05")
	return t.Format("2-Jan-2006")
}
