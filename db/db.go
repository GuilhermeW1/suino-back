package db

import (
	"fmt"
	"log"
	"os"

	"github.com/GuilhermeW1/backend-suino/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {

	dsn := os.Getenv("DB_CONNECTION")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	//testes
	if os.Getenv("SKIP_MIGRATIONS") == "true" {
		log.Println("SKIP_MIGRATIONS=true — pulando reset e AutoMigrate")
	} else {
		if os.Getenv("DB_RESET") == "true" && os.Getenv("ENV") != "prod" {
			resetDatabase(database)
		}
	}

	err = database.AutoMigrate(
		&model.Sow{},
		&model.Event{},
		&model.Cycle{},
	)

	if err != nil {
		log.Println("Erro ao rodar migrations", err)
	}

	fmt.Println("Conexao com o banco de dados concluida")
	DB = database
	return database
}

func resetDatabase(db *gorm.DB) {
	log.Println("DB_RESET=true — apagando e recriando o schema")
	db.Exec("DROP SCHEMA public CASCADE")
	db.Exec("CREATE SCHEMA public")
}
