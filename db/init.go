package db

import (
	"employee/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var EngineDB *gorm.DB

func Init() error {
	dsn := "host=localhost user=postgres dbname=opn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return err
	}
	EngineDB = db
	log.Log.Debug("Init Database Success")
	return nil
}