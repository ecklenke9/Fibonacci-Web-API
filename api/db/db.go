package db

import (
	"Fibonacci-Web-API/api/model"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
	"strconv"
	"strings"
)

// create var Postgres which will be used to store db connection
var postgresDB *sql.DB

func New() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// initialize the connection pool.
	connection, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	postgresDB = connection
	return
}

func GetMemoizedResults(fibNum int) int {

	var memoizedResults int

	// create the select sql query
	sqlStatement := `SELECT COUNT(*) FROM fibonacci WHERE fibnum <= $1`

	// execute the sql statement
	_ = postgresDB.QueryRow(sqlStatement, fibNum).Scan(&memoizedResults)

	// return empty fibonacci on error
	return memoizedResults
}

// insert one fibonacci in the db
func UpsertFibonacci(fibMap map[int]int) {
	// create the insert sql query
	// returning fibonacci number will return the ordinal of the inserted fibonacci
	beginStatement := `INSERT INTO fibonacci (ordinal, fibNum) VALUES`
	endStatement := ` ON CONFLICT (ordinal) DO NOTHING;`

	// Scan function will save the insert ordinal in the ordinal
	strb := strings.Builder{}
	strb.WriteString(beginStatement)

	fibs := make([]string, 0)
	for o, f := range fibMap {
		fibs = append(fibs, "("+strconv.Itoa(o)+","+strconv.Itoa(f)+"),")
	}

	for i, v := range fibs {
		if i+1 >= len(fibs) {
			v = strings.TrimRight(v, ",")
		}
		strb.WriteString(v)
	}

	strb.WriteString(endStatement)
	// execute the sql statement
	rows, err := postgresDB.Query(strb.String())
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}

	log.Printf("Fibonacci values by ordinal stored in SQL: %v", fibs)
	err = rows.Close()
	if err != nil {
		log.Printf("Unable to close rows: %v", err)
	}

	return
}

// get one fibonacci number from the db by its ordinal
func GetFibonacci(ordinal int) (model.Fibonacci, error) {
	// create a fibonacci of models.Fibonacci type
	var fibonacci model.Fibonacci

	// create the select sql query
	sqlStatement := `SELECT * FROM fibonacci WHERE ordinal=$1`

	// execute the sql statement
	row := postgresDB.QueryRow(sqlStatement, ordinal)

	// unmarshal the row object to fibonacci
	err := row.Scan(&fibonacci.Ordinal, &fibonacci.FibNum)

	switch err {
	// nil case is hit when ordinal is found in db
	case nil:
		return fibonacci, nil
	// sql.ErrNoRows case is hit when the ordinal is not in the db
	case sql.ErrNoRows:
		fibonacci.Ordinal = ordinal
		fibonacci.FibNum = fibonacciCalculation(ordinal)
		return fibonacci, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty fibonacci on error
	return fibonacci, err
}

// get all fibonacci numbers from the db
func GetAllFibonacci() ([]model.Fibonacci, error) {
	var fibonaccis []model.Fibonacci

	// create the select sql query
	sqlStatement := `SELECT * FROM fibonacci`

	// execute the sql statement
	rows, err := postgresDB.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var fibonacci model.Fibonacci

		// unmarshal the row object to fibonacci
		err = rows.Scan(&fibonacci.Ordinal, &fibonacci.FibNum)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the fibonacci in the fibonaccis slice
		fibonaccis = append(fibonaccis, fibonacci)
	}

	// return empty fibonacci on error
	return fibonaccis, err
}

// delete all fibonacci numbers in the db
func DeleteAllFibonacci() int {
	// create the delete sql query
	sqlStatement := `DELETE FROM fibonacci;`

	// execute the sql statement
	res, err := postgresDB.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return int(rowsAffected)
}

//------------------------- helper functions ----------------

func fibonacciCalculation(ordinal int) int {
	fn := make(map[int]int)
	for i := 0; i <= ordinal; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}

	UpsertFibonacci(fn)
	return fn[ordinal]
}
