package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBSource  string
	Port      string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not loaded, using system env")
	}

	return &Config{
		DBSource:  os.Getenv("DB_SOURCE"),
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}	
}
