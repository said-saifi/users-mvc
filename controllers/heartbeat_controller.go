package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Heartbeat(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "1.0.0",
		"server_time": time.Now(),
		"message":"running",
	})
}
