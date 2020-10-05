package controllers

import "github.com/gin-gonic/gin"

// Ping to health check server
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
