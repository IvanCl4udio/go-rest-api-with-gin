package database

import (
	"log"

	"github.com/ivancl4udio/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	// Cria automaticamente uma tabela no banco de dados baseado na struct que foi definida e passada como ponteiro
	DB.AutoMigrate(&models.Aluno{})
}
