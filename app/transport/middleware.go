package transport

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgEmptyAuthHeader})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgInvalidAuthHeader})

		return
	}

	userId, err := h.usecases.UserUseCase.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgProblemWithParseToken, Details: err.Error()})

		return
	}
	//возможно, нужно проверить, рабочий ли токен не выполнен ли по нему лог-аут
	// передав через ЮзКейс роспарсенный ИД в репозиторий и проверив его статус в БД

	c.Set(userCtx, userId)
}
