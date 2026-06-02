package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/adityadav2/eduflow-ai/backend/internal/database"
)

type Handler struct {
	mongoClient *mongo.Client
}

func NewHandler(mongoClient *mongo.Client) *Handler {
	return &Handler{
		mongoClient: mongoClient,
	}
}

func (h *Handler) Live(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  "alive",
		Service: "eduflow-api",
	})
}

func (h *Handler) Ready(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := database.PingMongo(ctx, h.mongoClient); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not_ready",
			"service": "eduflow-api",
			"mongodb": "down",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"service": "eduflow-api",
		"mongodb": "connected",
	})
}

func (h *Handler) DB(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := database.PingMongo(ctx, h.mongoClient); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"mongodb": "down",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"mongodb": "connected",
	})
}
