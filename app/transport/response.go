package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func newResponse(c *gin.Context, statusCode int, data any) {
	if data == nil {
		if statusCode != http.StatusNoContent {
			logrus.Error("Invalid data, expected nil")
		}

		return
	}
	c.AbortWithStatusJSON(statusCode, data)
}
