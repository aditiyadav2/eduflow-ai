package server

import (
	"github.com/adityadav2/eduflow-ai/backend/internal/health"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthCheck", health.Handler)
	return router
}
