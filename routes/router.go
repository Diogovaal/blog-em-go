package routes

import (
	"blog-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.ListarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuario)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)

	//login rota

	r.POST("/login", controllers.Login)
}
