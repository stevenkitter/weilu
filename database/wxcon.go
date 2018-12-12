package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

const (
	//WXUSER mysql user name
	WXUSER = "wx"
	//WXPASSWORD mysql password
	WXPASSWORD = "AxpYFTrjmlmjb1wq"
)

//WXDB sms db
func WXDB() (*gorm.DB, error) {
	db, err := ConnectDB(WXUSER, WXPASSWORD, WXUSER)
	if err != nil {
		log.Printf("connect db err : %v", err)
		return nil, err
	}
	go migrate(db)
	return db, err
}

func migrate(db *gorm.DB) {
	// db.AutoMigrate(&SMS{})
}
