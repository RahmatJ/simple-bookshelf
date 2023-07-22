package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) CheckHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
