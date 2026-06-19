package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id":   c.GetString("user_id"),
		"email":     c.GetString("email"),
		"role":      c.GetString("role"),
		"tenant_id": c.GetString("tenant_id"),
	})
}
