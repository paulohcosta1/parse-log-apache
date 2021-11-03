package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DbHost     = ""
	DbPort     = ""
	DbUser     = ""
	DbPassword = ""
	DbName     = ""
)

func init() {
	godotenv.Load("./.env")
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}
	godotenv.Load(".env." + env)

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
}
