package main

import (
	"Fibonacci-Web-API/pkg/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
