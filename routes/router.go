package routes

import (
	"blog-api/controllers"
	"blog-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	//publico
	r.POST("/usuarios", controllers.CriarUsuario)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)

	//crud

	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuario)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)

	//protegido
	admin := r.Group("/api/admin")
	admin.Use(middlewares.AuthMiddleware())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Área protegida ok"})
		})
	}

}
