package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// captura o header authorization

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token não fornecido"})
			c.Abort()
			return
		}
		// Esperado formato: "Bear <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mal formatado"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		secret := os.Getenv("SECRET_JWT")

		fmt.Println("Token:", tokenString)
		fmt.Println("SECRET_JWT:", secret)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//confirma que o algoritmo usado é o HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metodo de assinatura inválido")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido ou expirado"})
			c.Abort()
			return
		}
		//token invalido segue para proxima etapa
		c.Next()

	}
}
