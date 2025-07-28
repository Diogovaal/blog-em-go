package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Titulo   string `json:"titulo"`
	Conteudo string `json:"conteudo"`
}

var posts []Post

func CriarPost(c *gin.Context) {
    var post Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
        return
    }

    posts = append(posts, post)
    c.JSON(http.StatusCreated, post)
}

func ListarPosts(c *gin.Context) {
    c.JSON(http.StatusOK, posts)
}