package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	api := router.Group("api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!!",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "7777"
	}
	fmt.Println("Starting server on port:", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	router.Run(":" + port)
}
