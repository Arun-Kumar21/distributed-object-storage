package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEncryptionKey () []byte {
	LoadEnv()
	key := os.Getenv("ENCRYPTION_KEY")

	if len(key) != 16 {
		fmt.Println(key)
		log.Fatal("Invalid AES key length: must be exactly 16 bytes")
	}
	return []byte(key)
}
