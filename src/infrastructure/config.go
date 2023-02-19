package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
