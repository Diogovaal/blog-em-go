package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/models"
	"blog-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//variaveis ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	database.Connect()

	database.DB.AutoMigrate(&models.Usuario{})

	r := gin.Default()
	r.Use(middlewares.Logger())

	routes.SetupRoutes(r)

	//usa4 porta env

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
