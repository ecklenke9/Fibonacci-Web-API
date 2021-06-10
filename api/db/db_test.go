package db_test

import (
	"Fibonacci-Web-API/api/model"
	"fmt"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var db *gorm.DB

var fib = &model.Fibonacci{
	Ordinal: 12,
	FibNum:  144,
}

func TestConnectDataBase(t *testing.T) {
	// Create a new pool for Docker containers
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Pull an image, create a container based on it and set all necessary parameters
	opts := dockertest.RunOptions{
		Repository:   "mdillon/postgis",
		Tag:          "latest",
		Env:          []string{"POSTGRES_PASSWORD=mysecretpassword"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5433"},
			},
		},
	}

	// Run the Docker container
	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Exponential retry to connect to database while it is booting
	if err := pool.Retry(func() error {
		databaseConnStr := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
		db, err = gorm.Open("postgres", databaseConnStr)
		if err != nil {
			log.Println("Database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}

		// Tests if database is reachable
		return db.DB().Ping()
	}); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	log.Println("Initialize test database...")
	initTestDatabase()

	// Delete the Docker container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

}

var sampleFibonacci = []model.Fibonacci{
	{
		Ordinal: 11,
		FibNum:  89,
	},
	{
		Ordinal: 12,
		FibNum:  144,
	},
}

func initTestDatabase() {
	db.AutoMigrate(&model.Fibonacci{})

	db.Save(&sampleFibonacci[0])
	db.Save(&sampleFibonacci[1])
}
