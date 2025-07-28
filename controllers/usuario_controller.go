package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

var usuarios []Usuario

func CriarUsuario(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"erro": "JSON invalido"})
		return
	}

	initializers.DB.Create(&usuario)
	c.JSON(200, usuario)
}

func ListarUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}
