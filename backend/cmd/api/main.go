package main

import (
	"log"

	"github.com/adityadav2/eduflow-ai/backend/internal/config"
	"github.com/adityadav2/eduflow-ai/backend/internal/database"
	"github.com/adityadav2/eduflow-ai/backend/internal/server"
)

func main() {
	cfg := config.Load()

	mongoClient, err := database.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Fatal("failed to connect to MongoDB:", err)
	}

	router := server.NewRouter(mongoClient)

	log.Println("EduFlow API server started on port", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
