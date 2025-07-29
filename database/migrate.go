package database

import "blog-api/models"

func MigrarBanco() {
	err := DB.AutoMigrate(&models.Usuario{})
	if err != nil {
		panic("Erro ao migrar banco:" + err.Error())
	}
}
