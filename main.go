package main

import (
	"Fibonacci-Web-API/api/controller"
	"Fibonacci-Web-API/api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDataBase()

	// create routes
	r.GET("/api/fibonacci/memoizedResults/:fibnum", controller.GetMemoizedResults)
	r.GET("/api/fibonacci/:ordinal", controller.GetFibonacci)
	r.GET("/api/fibonacci/all", controller.GetAllFibonacci)
	r.DELETE("/api/fibonacci/clear", controller.DeleteAllFibonacci)

	r.Run(":8080")
}
