package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	response := Reponse{
		Status:  "ok",
		Service: "eduflow-api",
	}
	c.JSON(http.StatusOK, response)
}
