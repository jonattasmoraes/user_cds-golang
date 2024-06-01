package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/app-go/internal/app/user/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
		return nil, err
	}

	// Obter a string de conexão do banco de dados (DSN)
	dsn := os.Getenv("DSN")
	if dsn == "" {
		err := fmt.Errorf("DSN environment variable is not set")
		fmt.Println(err)
		return nil, err
	}

	// Conectar ao banco de dados
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error connecting to database: %v\n", err)
		return nil, err
	}

	err = database.AutoMigrate(models.User{})
	if err != nil {
		fmt.Printf("error migrating database: %v\n", err)
		return nil, err
	}

	// Retornar a conexão com o banco de dados
	return database, nil
}
