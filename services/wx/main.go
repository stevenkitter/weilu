package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stevenkitter/weilu/database"
	"github.com/stevenkitter/weilu/services/wx/server"
)

const port = ":51001"

func main() {
	db, err := database.WXDB()
	if err != nil {
		log.Panicf("wx server can not connect the mysql err : %v", err)
	}
	server := &server.Server{
		DB: db,
	}
	log.Printf("wx server is running in port %s", port)
	if err := server.Run(port); err != nil {
		log.Panicf("wx server run something is wrong err : %v", err)
	}
}
