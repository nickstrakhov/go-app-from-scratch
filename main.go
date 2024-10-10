package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.Run()
}
