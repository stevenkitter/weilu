package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stevenkitter/weilu/database"
	"github.com/stevenkitter/weilu/services/wx/server"
)

const port = ":51001"

// env
// @params WX_MYSQL_USER
// @params WX_MYSQL_PASSWORD
// @params WX_DATABASE
// @params WXAppSecret
func main() {
	db, err := database.WXDB()
	if err != nil {
		log.Panicf("wx server can not connect the mysql err : %v", err)
	}
	sv := &server.Server{
		DB: db,
	}
	log.Printf("wx server is running in port %s", port)
	if err := sv.Run(port); err != nil {
		log.Panicf("wx server run something is wrong err : %v", err)
	}
}
