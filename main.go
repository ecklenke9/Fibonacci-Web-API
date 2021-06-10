package main

import (
	"Fibonacci-Web-API/api/controller"
	"Fibonacci-Web-API/api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// set the router as the default provided by Gin
	r := gin.Default()

	// connect to Postgres db
	db.ConnectDataBase()

	// create routes
	r.GET("/api/fibonacci/memoizedResults/:fibnum", controller.GetMemoizedResults)
	r.GET("/api/fibonacci/:ordinal", controller.GetFibonacci)
	r.GET("/api/fibonacci/all", controller.GetAllFibonacci)
	r.DELETE("/api/fibonacci/clear", controller.DeleteAllFibonacci)

	// start serving the application
	r.Run(":8080")
}
