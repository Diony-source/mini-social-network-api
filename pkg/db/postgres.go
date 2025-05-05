package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dataSource string) *sql.DB {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("❌ failed to open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ failed to ping db: %v", err)
	}

	fmt.Println("✅ Connected to Postgres")
	return db
}
