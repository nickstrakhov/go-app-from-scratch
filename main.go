package main

import (
	"net/http"

	"go-app-from-scratch/config"
	"go-app-from-scratch/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []models.Users
		config.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.Run()
}
