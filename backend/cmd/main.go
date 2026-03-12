package main

import (
	"net/http"

	"github.com/Dhorq/team-task-management/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// router.GET("/task", getTask)

	router.Run(":5000")
}
