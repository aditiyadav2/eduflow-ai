package server

import (
	"github.com/adityadav2/eduflow-ai/backend/internal/health"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	healthHandler := health.NewHandler(mongoClient)

	router.GET("/health/live", healthHandler.Live)
	router.GET("/health/ready", healthHandler.Ready)
	router.GET("/health/db", healthHandler.DB)

	return router
}
