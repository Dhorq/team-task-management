package main

import (
	"net/http"

	"github.com/Dhorq/team-task-management/models"
	"github.com/Dhorq/team-task-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()
	models.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.SetupRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":5000")
}
