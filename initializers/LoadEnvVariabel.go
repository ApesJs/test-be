package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariable() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading Env File.")
	}
}
