package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
