package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var body struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	//buscar usuario por email
	var usuario models.Usuario
	if result := database.DB.Where("email = ?", body.Email).First(&usuario); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuario não encontrado"})
		return
	}

	//comparar a senha
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(body.Senha)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Senha inválida"})
		return
	}

	//criar token jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": usuario.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	//chave secreta
	secret := []byte("minha_env")

	tokenString, err := token.SignedString(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
