package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/my-gin-app/config"
	"github.com/yourname/my-gin-app/database"
	"github.com/yourname/my-gin-app/routes"
)

func main() {
	config.LoadEnv()
	database.ConnectMongoDB()

	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	routes.UserRoutes(r)

	r.Run(":8080")
}
