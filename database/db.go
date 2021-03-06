package database

import (
	"log"

	"github.com/Dagurasu56/go-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
