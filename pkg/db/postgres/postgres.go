package postgres

import (
	"Fibonacci-Web-API/pkg/model"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
)

// create connection with postgres db
func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// insert one fibonacci in the DB
func InsertFibonacci(fibonacci model.Fibonacci) int64 {

	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning fibonacci number will return the fib of the inserted fibonacci
	sqlStatement := `INSERT INTO fibonacci (fib, fibN) VALUES ($1, $2) RETURNING fib`

	// the inserted fib will store in this fib
	var fib int64

	// execute the sql statement
	// Scan function will save the insert fib in the fib
	err := db.QueryRow(sqlStatement, fibonacci.Fib, fibonacci.FibN).Scan(&fib)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", fib)

	// return the inserted fib
	return fib
}

// get one fibonacci from the DB by its fibonacci number
func GetFibonacci(fib int64) (model.Fibonacci, error) {
	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create a fibonacci of models.Fibonacci type
	var fibonacci model.Fibonacci

	// create the select sql query
	sqlStatement := `SELECT * FROM fibonacci WHERE fib=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, fib)

	// unmarshal the row object to fibonacci
	err := row.Scan(&fibonacci.Fib, &fibonacci.FibN)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return fibonacci, nil
	case nil:
		return fibonacci, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty fibonacci on error
	return fibonacci, err
}

// get one fibonacci from the DB by its fibonacci number
func GetAllFibonaccis() ([]model.Fibonacci, error) {
	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	var fibonaccis []model.Fibonacci

	// create the select sql query
	sqlStatement := `SELECT * FROM fibonacci`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var fibonacci model.Fibonacci

		// unmarshal the row object to fibonacci
		err = rows.Scan(&fibonacci.Fib, &fibonacci.FibN)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the fibonacci in the fibonaccis slice
		fibonaccis = append(fibonaccis, fibonacci)

	}

	// return empty fibonacci on error
	return fibonaccis, err
}

// update fibonacci in the DB
func UpdateFibonacci(fib int64, fibonacci model.Fibonacci) int64 {

	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE fibonacci SET fib=$2, fibN=$3 WHERE fib=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, fib, fibonacci.Fib, fibonacci.FibN)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete fibonacci in the DB
func DeleteFibonacci(fib int64) int64 {

	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM fibonacci WHERE fib=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, fib)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
