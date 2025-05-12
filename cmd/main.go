package main

import (
	"log"
	stdhttp "net/http"

	"mini-social-network-api/config"
	apphttp "mini-social-network-api/internal/http"
	"mini-social-network-api/pkg/db"
	"mini-social-network-api/pkg/logger"
)

func main() {
	logger.InitLogger()

	cfg := config.LoadConfig()
	database := db.ConnectPostgres(cfg)
	router := apphttp.NewRouter(cfg, database)

	logger.Log.WithField("port", cfg.Port).Info("Starting server.")
	log.Fatal(stdhttp.ListenAndServe(":"+cfg.Port, router))
}
