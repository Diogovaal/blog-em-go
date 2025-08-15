package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"blog-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//DTOs: estruturas para entrada/saída sem expor campos indevidos

type UsuarioInput struct {
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6"`
}

type UsuarioUpdateInput struct {
	Nome  *string `json:"nome,omitempty"`
	Email *string `json:"email,omitempty"`
	Senha *string `json:"senha,omitempty"`
}

type UsuarioOut struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Helper para “sanitizar” a saída (nunca expor senha)
func toUsuarioOut(u models.Usuario) UsuarioOut {
	return UsuarioOut{ID: u.ID, Nome: u.Nome, Email: u.Email}
}

func CriarUsuario(c *gin.Context) {
	var in UsuarioInput

	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	//gerar hash da senha antes de persistir
	hash, err := utils.HashPassword(in.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao processar senha"})
		return
	}

	u := models.Usuario{
		Nome:  in.Nome,
		Email: in.Email,
		Senha: hash, //guarda só o hash
	}
	if err := database.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toUsuarioOut(u))

}

func ListarUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	if err := database.DB.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	out := make([]UsuarioOut, 0, len(usuarios))
	for _, u := range usuarios {
		out = append(out, toUsuarioOut(u))
	}
	c.JSON(http.StatusOK, out)
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")
	var u models.Usuario

	if result := database.DB.First(&u, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, toUsuarioOut(u))
}

func AtualizarUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var in UsuarioUpdateInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var u models.Usuario

	if err := database.DB.First(&u, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuario não encontrado"})
		return
	}

	//Atualizar os campos recebidos no JSON
	if in.Nome != nil {
		u.Nome = *in.Nome
	}
	if in.Email != nil {
		u.Email = *in.Email
	}

	if in.Senha != nil && *in.Senha != "" {
		hash, err := utils.HashPassword(*in.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao processar senha"})
			return
		}
		u.Senha = hash
	}

	if err := database.DB.Save(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toUsuarioOut(u))

}

//delete

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Usuario{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
