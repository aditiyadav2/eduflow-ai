package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func healthHandler(c *gin.Context) {
	response := HealthResponse{
		Status:  "ok",
		Service: "eduflow-api",
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.GET("/health", healthHandler)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
