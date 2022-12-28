package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func newResponse(c *gin.Context, statusCode int, data any) {
	if data == nil {
		if statusCode != http.StatusNoContent {
			log.Printf("Invalid data, expected nil")
		}

		return
	}
	c.AbortWithStatusJSON(statusCode, data)
}
