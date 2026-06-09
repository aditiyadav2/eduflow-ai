package server

import (
	"github.com/adityadav2/eduflow-ai/backend/internal/auth"
	"github.com/adityadav2/eduflow-ai/backend/internal/database"
	"github.com/adityadav2/eduflow-ai/backend/internal/health"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Database
	db := database.GetDatabase(mongoClient)

	// Health module
	healthHandler := health.NewHandler(mongoClient)

	router.GET("/health/live", healthHandler.Live)
	router.GET("/health/ready", healthHandler.Ready)
	router.GET("/health/db", healthHandler.DB)

	// Auth module
	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)

	api := router.Group("/api/v1")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", authHandler.Register)
		}
	}

	return router
}
