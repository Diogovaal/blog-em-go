package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var usuario models.Usuario

func CriarUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	result := database.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar no banco"})

	}

	c.JSON(http.StatusOK, usuario)
}

func ListarUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuario)
}
