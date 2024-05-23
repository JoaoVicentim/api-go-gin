package database

import (
	"log"

	"github.com/JoaoVicentim/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	//com ele, podemos inserir qual struct queremos colocar la no banco de dados (&models.Aluno{}), la no postgres
	DB.AutoMigrate(&models.Aluno{})
}
