package transport

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userPermission      = "userRole"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgEmptyAuthHeader})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgInvalidAuthHeader})

		return
	}

	userId, userRole, err := h.usecases.UserUseCase.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgProblemWithParseToken, Details: err.Error()})

		return
	}

	c.Set(userCtx, userId)
	c.Set(userPermission, userRole)
}
