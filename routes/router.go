package routes

import (
	"blog-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	usuarioRoutes := r.Group("/usuarios")
	{
		usuarioRoutes.POST("/", controllers.CriarUsuario)
		usuarioRoutes.GET("/", controllers.ListarUsuarios)
	}

	postRoutes := r.Group("/posts")
	{
		postRoutes.POST("/", controllers.CriarPost)
		postRoutes.GET("/", controllers.ListarPosts)
	}
}
