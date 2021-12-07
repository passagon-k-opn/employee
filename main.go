package main

import (
	"employee/db"
	"employee/log"
	"employee/router"
)

func main() {
	log.InitLogrus()
	err := db.Init()
	if err !=nil{
		return
	}
	router.InitRouter()

}
