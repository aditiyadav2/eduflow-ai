package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "welcome to admin dashboard",
		"user_id":   c.GetString("user_id"),
		"email":     c.GetString("email"),
		"role":      c.GetString("role"),
		"tenant_id": c.GetString("tenant_id"),
	})
}
