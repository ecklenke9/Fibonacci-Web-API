package router

import (
	"Fibonacci-Web-API/api/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// response format
type response struct {
	Ordinal int    `json:"ordinal,omitempty"`
	Message string `json:"message,omitempty"`
}

// New is exported and used in main.go
func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/fibonacci/memoizedResults/{fibnum}", getMemoizedResults).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fibonacci/{ordinal}", getFibonacci).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fibonacci/all", getAllFibonacci).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fibonacci/clear", deleteAllFibonacci).Methods("DELETE", "OPTIONS")

	return router
}

// getMemoizedResults will return the number of memoized results less than the given value
// (e.g. there are 12 intermediate results less than 120)
func getMemoizedResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get the "value" from the request params
	params := mux.Vars(r)

	// convert the value type from string to int
	value, err := strconv.Atoi(params["fibnum"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the GetMemoizedResults function with value to retrieve the number of memoized results
	memoizedResults := db.GetMemoizedResults(value)

	// send the response
	var msg = "there are %d intermediate results less than %d"
	json.NewEncoder(w).Encode(response{Message: fmt.Sprintf(msg, memoizedResults, value)})
}

// getFibonacci will return a single fibonacci by its ordinal
func getFibonacci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get the "ordinal" from the request params
	params := mux.Vars(r)

	// convert the ordinal type from string to int
	ordinal, err := strconv.Atoi(params["ordinal"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getFibonacci function with ordinal to retrieve a single fibonacci
	fibonacci, err := db.GetFibonacci(int(ordinal))
	if err != nil {
		log.Fatalf("Unable to get fibonacci number. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(fibonacci)
}

// getAllFibonacci will return all the fibonacci numbers stored in db
func getAllFibonacci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the fibonacci numbers in the db
	fibonaccis, err := db.GetAllFibonacci()

	if err != nil {
		log.Fatalf("Unable to get all fibonacci numbers. %v", err)
	}

	// send all the fibonacci numbers as response
	json.NewEncoder(w).Encode(fibonaccis)
}

// deleteAllFibonacci will delete all fibonacci numbers in the postgres db
func deleteAllFibonacci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// call the deleteAllFibonacci, convert the int to int
	deletedRows := db.DeleteAllFibonacci()

	// format the message string
	msg := fmt.Sprintf("Fibonacci numbers successfully deleted. Total rows/record affected %v", deletedRows)

	// format the response message
	res := response{
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
