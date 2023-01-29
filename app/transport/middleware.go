package transport

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userPermission      = "userRole"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logrus.Warn(MsgEmptyAuthHeader)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgEmptyAuthHeader})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		logrus.Warn(MsgInvalidAuthHeader)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgInvalidAuthHeader})

		return
	}

	userId, userRole, err := h.usecases.UserUseCase.ParseToken(headerParts[1])
	if err != nil {
		logrus.Warn(err)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgProblemWithParseToken, Details: err.Error()})

		return
	}
	logrus.Debugf("Received token from user: %v. Token is correct", userId)
	c.Set(userCtx, userId)
	c.Set(userPermission, userRole)
}
