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
	authorHandler := handler.NewAuthorHandler()
	bookHandler := handler.NewBookHandler()

	// set route
	r.GET("/ping", healthCheckHandler.CheckHealth)

	r.GET("/authors/:id", authorHandler.GetAuthorById)

	r.GET("/books", bookHandler.GetBook)
	r.POST("/books", bookHandler.CreateBook)
	r.PATCH("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

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
