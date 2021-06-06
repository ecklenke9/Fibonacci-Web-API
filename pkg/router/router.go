package router

import (
	"Fibonacci-Web-API/pkg/db/postgres"
	"Fibonacci-Web-API/pkg/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// response format
type response struct {
	Fib     int64  `json:"fib,omitempty"`
	Message string `json:"message,omitempty"`
}

// New is exported and used in main.go
func New() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/fibonacci/{fib}", GetFibonacci).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fibonacci", GetAllFibonacci).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newFibonacci", CreateFibonacci).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/fibonacci/{fib}", UpdateFibonacci).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteFibonacci/{fib}", DeleteFibonacci).Methods("DELETE", "OPTIONS")

	return router
}

// CreateFibonacci creates a fibonacci in the postgres db
func CreateFibonacci(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty fibonacci of type model.Fibonacci
	var fibonacci model.Fibonacci

	// decode the json request to fibonacci
	err := json.NewDecoder(r.Body).Decode(&fibonacci)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insertFibonacci function and pass the fibonacci number
	insertFib := postgres.InsertFibonacci(fibonacci)

	// format a response object
	res := response{
		Fib:     insertFib,
		Message: "Successfully inserted Fibonacci number",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetFibonacci will return a single fibonacci by its fibonacci number
func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the fibonacci number from the request params, key is "fib"
	params := mux.Vars(r)

	// convert the fib type from string to int
	fib, err := strconv.Atoi(params["fib"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getFibonacci function with fibonacci number to retrieve a single fibonacci
	fibonacci, err := postgres.GetFibonacci(int64(fib))

	if err != nil {
		log.Fatalf("Unable to get fibonacci. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(fibonacci)
}

// GetAllFibonacci will return all the fibonaccis
func GetAllFibonacci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the fibonaccis in the db
	fibonaccis, err := postgres.GetAllFibonaccis()

	if err != nil {
		log.Fatalf("Unable to get all fibonacci. %v", err)
	}

	// send all the fibonaccis as response
	json.NewEncoder(w).Encode(fibonaccis)
}

// UpdateFibonacci update fibonacci's detail in the postgres db
func UpdateFibonacci(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the fibonacci number from the request params, key is "fib"
	params := mux.Vars(r)

	// convert the fib type from string to int
	fib, err := strconv.Atoi(params["fib"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty fibonacci of type models.Fibonacci
	var fibonacci model.Fibonacci

	// decode the json request to fibonacci
	err = json.NewDecoder(r.Body).Decode(&fibonacci)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update fibonacci to update the fibonacci
	updatedRows := postgres.UpdateFibonacci(int64(fib), fibonacci)

	// format the message string
	msg := fmt.Sprintf("Fibonacci updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		Fib:     int64(fib),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteFibonacci delete fibonacci's detail in the postgres db
func DeleteFibonacci(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the fibonacci number from the request params, key is "fib"
	params := mux.Vars(r)

	// convert the fib in string to int
	fib, err := strconv.Atoi(params["fib"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteFibonacci, convert the int to int64
	deletedRows := postgres.DeleteFibonacci(int64(fib))

	// format the message string
	msg := fmt.Sprintf("Fibonacci updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		Fib:     int64(fib),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
