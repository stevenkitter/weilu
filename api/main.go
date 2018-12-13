package main

import (
	"log"

	"github.com/stevenkitter/weilu/api/manager"
)

//Port every api start from 8100
const Port = ":8100"

func main() {
	m := manager.NewManager()
	if err := m.Run(Port); err != nil {
		log.Panicf("weilu api run err: %v", err)
	}
}
