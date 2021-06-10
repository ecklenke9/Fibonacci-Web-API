package db

import (
	"Fibonacci-Web-API/api/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
)

// create var Postgres which will be used to store db connection
var Postgres *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	// set up environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSLMODE")

	// initialize a new db connection with the environment variables
	database, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, sslMode, dbPassword))
	if err != nil {
		panic("failed to connect to database")
	}

	// create table 'fibonaccis' if it does not currently exist
	database.AutoMigrate(&model.Fibonacci{})

	// assign the newly created 'database' to the global variable, 'Postgres'
	Postgres = database
}
