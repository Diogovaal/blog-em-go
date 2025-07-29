package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/models"
	"blog-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	database.DB.AutoMigrate(&models.Usuario{})
	database.MigrarBanco()

	r := gin.Default()
	r.Use(middlewares.Logger())

	routes.SetupRoutes(r)

	r.Run(":8080")
}
