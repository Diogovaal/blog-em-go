package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
