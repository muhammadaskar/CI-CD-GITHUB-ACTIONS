package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Hello World!!!"}`, w.Body.String())
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!!",
		})
	})
	return router
}
