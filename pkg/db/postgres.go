package db

import (
	"database/sql"
	"fmt"
	"log"
	"mini-social-network-api/config"

	_ "github.com/lib/pq"
)

func getDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
}

func ConnectPostgres(cfg *config.Config) *sql.DB {
	dsn := getDSN(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ failed to open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ failed to ping db: %v", err)
	}

	fmt.Println("✅ Connected to Postgres")
	return db
}