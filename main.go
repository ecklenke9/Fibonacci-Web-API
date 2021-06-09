package main

import (
	"Fibonacci-Web-API/api/db"
	"Fibonacci-Web-API/api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.New()

	r := router.New()
	fmt.Println("Starting server on the port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
