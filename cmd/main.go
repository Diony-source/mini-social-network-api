package main

import (
	"log"
	stdhttp "net/http"

	"mini-social-network-api/config"
	apphttp "mini-social-network-api/internal/http"
	"mini-social-network-api/pkg/db"
)

func main() {
	cfg := config.LoadConfig()
	database := db.ConnectPostgres(cfg.DBSource)
	router := apphttp.NewRouter(cfg, database)

	log.Printf("🚀 Server running on port %s", cfg.Port)
	log.Fatal(stdhttp.ListenAndServe(":"+cfg.Port, router))
}
