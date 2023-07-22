package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-bookshelf/cmd/common"
	"simple-bookshelf/cmd/handler"
)

func SetupRouter() (*gin.Engine, error) {
	//	init gin
	r := gin.Default()

	// Setup handler
	healthCheckHandler := handler.NewHealthCheckHandler()

	// set route
	r.GET("/ping", healthCheckHandler.CheckHealth)

	// run gin
	err := r.Run()
	return r, err
}

func main() {
	// connect to db
	common.NewConnection()

	//	init gin
	_, err := SetupRouter()

	if err != nil {
		log.Fatal(err)
		return
	}
}
