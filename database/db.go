package database

import (
	"fmt"
	"log"

	"github.com/Ulpio/gin-api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connectString := "host=localhost user=root password=root dbname=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
	fmt.Println("Conexão com banco de dados estabelecida")
}
