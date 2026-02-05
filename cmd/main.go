package main

import (
	"employee-management/config"
	"employee-management/routers"
	"employee-management/internal/db"
	"log"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	database, err := db.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Apply migrations
	db.MigrateDB(database)

	router := routers.SetupRouter(cfg, database)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
