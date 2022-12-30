package transport

import (
	"net/http"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var input entities.User

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddUser(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type LogInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) LogIn(c *gin.Context) {
	var input LogInInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	token, err := h.usecases.GenerateToken(input.Login, input.Password)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) LogOut(c *gin.Context) {

}

func (h *Handler) CreateFilm(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GetAllFilms(c *gin.Context) {

}

func (h *Handler) GetFilmByID(c *gin.Context) {

}

func (h *Handler) AddToFavourite(c *gin.Context) {

}

func (h *Handler) AddToWishlist(c *gin.Context) {

}

func (h *Handler) GetCSVFile(c *gin.Context) {

}
