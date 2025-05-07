package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	JWTSecret  string
}

func LoadConfig() *Config {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBPassword == "" ||
		cfg.DBName == "" || cfg.Port == "" || cfg.JWTSecret == "" {
		log.Fatal("❌ Missing required environment variables")
	}

	return cfg
}
