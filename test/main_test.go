package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simple-bookshelf/cmd/handler"
	"testing"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestPingPong(t *testing.T) {
	// Given
	expectedResult := `{"message":"pong!"}`
	r := SetupRouter()

	// Setup Handler
	healthCheckHandler := handler.NewHealthCheckHandler()

	r.GET("/ping", healthCheckHandler.CheckHealth)

	// When
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedResult, w.Body.String())
}
