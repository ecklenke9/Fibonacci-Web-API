package db

import (
	"Fibonacci-Web-API/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgres golang driver
)

// create var Postgres which will be used to store db connection
var Postgres *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("postgres", "host=postgres-emma-8612.aivencloud.com port=15741 user=avnadmin dbname=defaultdb sslmode=require password=ohx89khk1wmr6d2o")
	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&model.Fibonacci{})
	Postgres = database
}
