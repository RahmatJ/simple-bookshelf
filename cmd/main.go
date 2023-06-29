package main

import (
	"github.com/gin-gonic/gin"
	"simple-bookshelf/cmd/common"
	"simple-bookshelf/cmd/controllers"
)

func SetupRouter() (*gin.Engine, error) {
	//	init gin
	r := gin.Default()
	r.GET("/ping", controllers.HealthCheckHandler)
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
		panic(err)
	}
}
