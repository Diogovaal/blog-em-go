package controllers

import (
	"blog-api/database"
	"blog-api/models"
	"blog-api/utils"
	"net/http"
	"os"

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
	accessToken, err := utils.GenerateAcessToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "não foi possivel gerar o token"})
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "não foi possivel gerar o token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token é obrigatório"})
		return
	}

	//validar o refresh token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_REFRESH")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}
	userID := uint(claims["user_id"].(float64))

	//gerar novos acess tokens
	newAcessToken, err := utils.GenerateAcessToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível gerar o token de acesso"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAcessToken,
	})
}
