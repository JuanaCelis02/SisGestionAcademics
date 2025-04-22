package main

import (
	"log"
	"os"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/api/routes"
	"uptc/sisgestion/pkg/config"
	"uptc/sisgestion/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.NewConfig()

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	router := gin.Default()

	router.Use(middlewares.Logger())
	router.Use(gin.Recovery())

	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := database.MigrateModels(db); err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}