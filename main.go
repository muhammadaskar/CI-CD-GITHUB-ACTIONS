package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	api := router.Group("api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":7777")

}
