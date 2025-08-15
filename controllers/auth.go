package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"blog-api/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

func Login(c *gin.Context) {

	var in LoginInput

	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	//buscar usuario por email
	var u models.Usuario
	if result := database.DB.Where("email = ?", in.Email).First(&u); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuario não encontrado"})
		return
	}

	//comparar a senha digitada com hash do banco
	if ok := utils.CheckPasswordHash(in.Senha, u.Senha); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "usuário ou senha inválidos"})
		return
	}
	//criar token jwt
	secret := os.Getenv("SECRET_JWT")
	claims := jwt.MapClaims{
		"sub":   u.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
		"email": u.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Register(c *gin.Context) {

}
