package controller

import (
	"Fibonacci-Web-API/api/db"
	"Fibonacci-Web-API/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// create var cache to store map[ordinal]fibnum
var cache map[int]int

func GetMemoizedResults(c *gin.Context) {
	// create var memoizedResults to store number of memoizedResults less than a given value
	var memoizedResults int
	fibNum := c.Param("fibnum")

	// reach out to Postgres to count the number of rows less than or equal to the fibNum given
	// SELECT count(*) FROM fibonacci WHERE fib_num <= $1;
	db.Postgres.Model(&model.Fibonacci{}).Where("fib_num <= $1", fibNum).Count(&memoizedResults)

	// return the intermediate results less than given fibNum
	c.JSON(http.StatusOK, gin.H{"memoizedResults": memoizedResults})
}

func GetFibonacci(c *gin.Context) {
	// create var fibResult to store the fibonacci results
	var fibResult model.Fibonacci
	// create var fibNum to store fibNum from fibResult
	var fibNum int
	// create var fibArray to store an array of Fibonacci
	var fibArray []model.Fibonacci

	// pull the ordinal from the request url and convert to int
	ordIntVal, err := strconv.Atoi(c.Param("ordinal"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// reach out to Postgres to find the ordinal
	// SELECT * FROM fibonacci WHERE ordinal = 'ordinal' ORDER BY id LIMIT 1;
	if err := db.Postgres.Where("ordinal = ?", c.Param("ordinal")).First(&fibResult).Error; err != nil {
		// if err do calculation for the ordinal given and all less than
		fibNum, fibArray = fibonacciCalculation(ordIntVal)
		insertFibonacci(fibArray)
	}

	// if the fibNum instantiated prior is zero, return fibNum from fibResult
	if fibNum == 0 {
		fibNum = fibResult.FibNum
	}

	// return the fibNum of the ordinal passed in
	c.JSON(http.StatusOK, gin.H{"fibonacciNumber": fibNum})
	return
}

func GetAllFibonacci(c *gin.Context) {
	// create var fibonacciArray to store an array of Fibonacci
	var fibArray []model.Fibonacci

	// reach out to Postgres and select all rows
	// SELECT * FROM fibonaccis;
	// unmarshall results from db into fibonacciArray
	db.Postgres.Find(&fibArray)

	// return the fibonacciArray
	c.JSON(http.StatusOK, gin.H{"allFibonacciResults": fibArray})
}

func DeleteAllFibonacci(c *gin.Context) {
	// create var fibonacci
	var fibonacci model.Fibonacci

	// clear cache
	cache = nil

	// reach out to Postgres and delete all rows
	// DELETE * FROM fibonaccis;
	// clear the database
	db.Postgres.Delete(&fibonacci)
	c.JSON(http.StatusOK, gin.H{"message": "database cleared"})
}

//------------------------- helper functions ----------------

func insertFibonacci(fibonacciArray []model.Fibonacci) {
	// insert all missing fibonacci numbers
	if cache == nil {
		cache = make(map[int]int, 0)
	}
	for _, fib := range fibonacciArray {
		// use ok idiom to determine if
		// ordinal has been processed already
		_, ok := cache[fib.Ordinal]
		if !ok {
			// ordinal has not been processed
			// create it in db
			cache[fib.Ordinal] = fib.FibNum
			db.Postgres.Create(&fib)
		}
	}

}

func fibonacciCalculation(ordinal int) (int, []model.Fibonacci) {
	// create map to hold ordinals and fibNums
	fn := make([]model.Fibonacci, 0)
	// creat var result to hold resulting fibNum
	var result int

	for i := 0; i <= ordinal; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1].FibNum + fn[i-2].FibNum
		}
		fn = append(fn, model.Fibonacci{
			Ordinal: i,
			FibNum:  f,
		})
		if i+1 >= ordinal {
			result = f
		}
	}

	return result, fn
}
