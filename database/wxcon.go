package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	//WXUSER mysql user name
	WXUSER = os.Getenv("WX_MYSQL_USER")
	//WXPASSWORD mysql password
	WXPASSWORD = os.Getenv("WX_MYSQL_PASSWORD")
	//WXDATABASE wx database
	WXDATABASE = os.Getenv("WX_DATABASE")
)

//WXDB sms db
func WXDB() (*gorm.DB, error) {
	db, err := ConnectDB(WXUSER, WXPASSWORD, WXDATABASE)
	if err != nil {
		log.Printf("connect db err : %v", err)
		return nil, err
	}
	go migrate(db)
	return db, err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&WXComponent{})
}
