package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// TODO: Rework
	// err := godotenv.Load(".env.local")
	// if err != nil {
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// } else if os.Getenv("ENV") == "prod" {
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// }
}
