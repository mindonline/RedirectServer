package Application

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetEnv(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	} else {
		return defaultValue
	}
}
