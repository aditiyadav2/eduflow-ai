package server

import (
	"github.com/adityadav2/eduflow-ai/backend/internal/auth"
	"github.com/adityadav2/eduflow-ai/backend/internal/database"
	"github.com/adityadav2/eduflow-ai/backend/internal/health"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	// middleware package intentionally not used here; remove import when adding middleware
	"github.com/adityadav2/eduflow-ai/backend/internal/admin"
	"github.com/adityadav2/eduflow-ai/backend/internal/middleware"
	"github.com/adityadav2/eduflow-ai/backend/internal/profile"
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
	authHandler := auth.NewHandler(authService, "eduflow-secret-key")

	api := router.Group("/api/v1")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", authHandler.Register)
			authRoutes.POST("/login", authHandler.Login)
		}
	}
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware("eduflow-secret-key"))
	{
		protected.GET("/profile", profile.GetProfile)
		adminRoutes := protected.Group("/admin")
		adminRoutes.Use(middleware.RequireRole("admin"))
		{
			adminRoutes.GET("/dashboard", admin.Dashboard)
		}

	}

	return router

}
