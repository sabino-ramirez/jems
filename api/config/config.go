package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Port  string
	DBurl string
}

func NewAppConfig() (App, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	appConfig := App{
		Port:  os.Getenv("PORT"),
		DBurl: os.Getenv("DB_URL"),
	}

	return appConfig, nil
}

func InitDB(dburl string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
