package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(name string) string {
	godotenv.Load(".env")
	switch name {
	case "DB_HOST":
		return os.Getenv(name)
	case "DB_PORT":
		return os.Getenv(name)
	case "DB_USER":
		return os.Getenv(name)
	case "DB_PASSWORD":
		return os.Getenv(name)
	case "DB_NAME":
		return os.Getenv(name)
	case "SERVER_PORT":
		return os.Getenv(name)
	case "JWT_SECRET":
		return os.Getenv(name)
	case "JWT_SECRET_ADMIN":
		return os.Getenv(name)
	default:
		panic(name + " doesnt exist in file .env")
	}
}
