package routes

import (
	"blog-api/controllers"
	"blog-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	v1 := r.Group("/api")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		//grupo protegido por jwt
		protected := v1.Group("/admin")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Bem-vindo a área protegida"})
			})
		}
	}

	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuario)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)

	//login rota

	r.POST("/login", controllers.Login)
}
