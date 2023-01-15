package transport

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var inputUserData entities.User

	if err := c.BindJSON(&inputUserData); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddUser(inputUserData)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
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

	token, err := h.usecases.GenerateAddToken(input.Login, input.Password)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) LogOut(c *gin.Context) {
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

	userId, _, err := h.usecases.UserUseCase.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgProblemWithParseToken, Details: err.Error()})

		return
	}

	if err := h.usecases.SignOut(userId, headerParts[1]); err != nil {
		newResponse(c, http.StatusNotFound, Response{Message: MsgNotFound, Details: err.Error()})

		return
	}
	c.JSON(http.StatusNoContent, nil)
	//log.Printf("logout is completed for user %s", userId)
}

func (h *Handler) CreateFilm(c *gin.Context) {
	var inputFilmData entities.Film

	role, _ := c.Get(userPermission)
	if role != "admin" {
		newResponse(c, http.StatusForbidden, Response{Message: MsgHaveNotPermission})

		return
	}

	if err := c.BindJSON(&inputFilmData); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	if err := h.usecases.ValidateFilmData(inputFilmData); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	directorId, err := h.usecases.GetDirectorId(inputFilmData)
	if err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddFilm(inputFilmData, directorId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"film_id": id,
	})
}

func (h *Handler) GetAllFilms(c *gin.Context) {

}

func (h *Handler) GetFilmByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgInvalidIDParam})
		return
	}

	film, err := h.usecases.GetFilmById(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, film)
}

type FilmList struct {
	Name string `json:"name"`
}

func (h *Handler) AddToFavourite(c *gin.Context) {
	userId, _ := c.Get(userCtx)

	var filmName FilmList
	if err := c.BindJSON(&filmName); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	filmID, err := h.usecases.GetFilmID(filmName.Name)
	if err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddFilmToFavourite(userId, filmID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id film in my favourite list list": id,
	})
}

func (h *Handler) AddToWishlist(c *gin.Context) {
	userId, _ := c.Get(userCtx)

	var filmName FilmList
	if err := c.BindJSON(&filmName); err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	filmID, err := h.usecases.GetFilmID(filmName.Name)
	if err != nil {
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddToWishlist(userId, filmID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id film in my wish list list": id,
	})
}

func (h *Handler) GetCSVFile(c *gin.Context) {

}
