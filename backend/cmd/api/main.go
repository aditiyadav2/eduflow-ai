package main

import (
	"log"

	"github.com/adityadav2/eduflow-ai/backend/internal/server"
)

func main() {
	router := server.NewRouter()
	log.Println("EduFlow API server started on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("server failed to start: ", err)
	}
}
