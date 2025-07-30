package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CriarUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	//criptografar a senha antes de salvar

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao criptografar a senha"})
		return
	}

	usuario.Senha = string(senhaHash)

	result := database.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return

	}

	c.JSON(http.StatusOK, usuario)
}

func ListarUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	database.DB.Find(&usuarios)
	c.JSON(http.StatusOK, usuarios)
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario

	if result := database.DB.First(&usuario, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario

	if result := database.DB.First(&usuario, id); result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"erro": "Usuario não encontrado"})
		return
	}

	//Atualizar os campos recebidos no JSON
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	database.DB.Save(&usuario)
	c.JSON(http.StatusOK, usuario)

}

//delete

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	if result := database.DB.Delete(&models.Usuario{}, id); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuario não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "usuario deletado com sucesso"})
}
