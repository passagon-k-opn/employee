package main

import (
	"employee/db"
	"employee/router"
)

func main() {

	err := db.Init()
	if err !=nil{
		return
	}
	router.InitRouter()

}
