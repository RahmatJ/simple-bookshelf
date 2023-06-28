package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-bookshelf/cmd/common"
)

func main() {
	// connect to db
	common.NewConnection()

	//	init gin
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	r.Run()
}
