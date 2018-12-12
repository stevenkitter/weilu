package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

const (
	//DEFURL default mysql url
	DEFURL = "115.159.222.199:3306"
)

//ConnectDB connect database
func ConnectDB(user, password, database string) (*gorm.DB, error) {
	path := os.Getenv("MYSQL_URL")
	if path == "" {
		path = DEFURL
	}
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, password, path, database)
	db, err := gorm.Open("mysql", sqlURL)
	return db, err
}
