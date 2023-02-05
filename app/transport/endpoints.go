package transport

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/globals"
	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/cfg"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var inputUserData entities.User

	if err := c.BindJSON(&inputUserData); err != nil {
		logrus.Warn(err)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	config, err := cfg.GetViperConfig()
	if err != nil {
		logrus.Errorf("Failed to get Viper config: %s", err)

		return
	}

	id, err := h.usecases.AddUser(inputUserData, config.Salt)
	if err != nil {
		if errors.Is(err, globals.ErrDuplicateLogin) {
			logrus.Warnf("Attempt to add user with an existing login: %v.", inputUserData.Login)
			newResponse(c, http.StatusConflict, Response{Message: MsgBadRequest, Details: err.Error()})

			return
		}
		if errors.Is(err, globals.ErrIncorrectUserData) {
			logrus.Warnf("Attempt to add user with incorect user data: %v.", inputUserData.Login)
			newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

			return
		}

		logrus.Errorf("Attempt to add user: %v.", err)
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalServerErr, Details: err.Error()})

		return
	}
	logrus.Infof("User successfully created. ID: %v", id)
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) LogIn(c *gin.Context) {
	var input entities.LogInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Warn(err)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	config, err := cfg.GetViperConfig()
	if err != nil {
		logrus.Errorf("Failed to get Viper config: %s", err)

		return
	}

	token, id, err := h.usecases.GenerateAddToken(input.Login, input.Password, config.SigningKey, config.Salt)
	if err != nil {
		if errors.Is(err, globals.ErrNotFound) {
			logrus.Warnf("Attempt to log in user: %v. %v.", input.Login, err)
			newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

			return
		}
		logrus.Errorf("Attempt to log in user: %v. %v.", input.Login, err)
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalServerErr, Details: err.Error()})

		return
	}

	logrus.Infof("User: %v ID: %v logging completed successfully, token sended to user", input.Login, id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"userID": id,
		"token":  token,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logrus.Warnf("Attempt to log out, %v ", MsgEmptyAuthHeader)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgEmptyAuthHeader})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logrus.Warnf("Attempt to log out, %v ", MsgInvalidAuthHeader)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgInvalidAuthHeader})

		return
	}

	config, err := cfg.GetViperConfig()
	if err != nil {
		logrus.Errorf("Failed to get Viper config: %s", err)

		return
	}

	userId, _, err := h.usecases.UserUseCase.ParseToken(headerParts[1], config.SigningKey)
	if err != nil {
		logrus.Warnf("Attempt to log out, %v ", err)
		newResponse(c, http.StatusUnauthorized, Response{Message: MsgProblemWithParseToken, Details: err.Error()})

		return
	}

	if err := h.usecases.SignOut(userId, headerParts[1]); err != nil {
		logrus.Warnf("Attempt to log out, %v ", err)
		newResponse(c, http.StatusNotFound, Response{Message: MsgNotFound, Details: err.Error()})

		return
	}
	logrus.Infof("User %v completed log out", userId)
	c.JSON(http.StatusNoContent, nil)

}

func (h *Handler) CreateFilm(c *gin.Context) {
	var inputFilmData entities.Film
	userId, _ := c.Get(userCtx)
	role, _ := c.Get(userPermission)
	if role != "admin" {
		logrus.Warnf("The user with id: %v is not administrator. Only administrator can add new movies", userId)
		newResponse(c, http.StatusForbidden, Response{Message: MsgHaveNotPermission})

		return
	}

	if err := c.BindJSON(&inputFilmData); err != nil {
		logrus.Errorf("Attempt to create movie: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	if err := h.usecases.ValidateFilmData(inputFilmData); err != nil {
		logrus.Warnf("Attempt to create movie vith wrong parametres: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	directorId, err := h.usecases.GetDirectorId(inputFilmData)
	if err != nil {
		logrus.Warnf("Attempt to create movie vith wrong parametres: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	id, err := h.usecases.AddFilm(inputFilmData, directorId)
	if err != nil {
		logrus.Warnf("Attempt to create movie vith wrong parametres: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalServerErr, Details: err.Error()})

		return
	}
	logrus.Infof("User: %v successfully created new film, with film ID: %v", userId, id)
	c.JSON(http.StatusCreated, map[string]interface{}{
		"film_id": id,
	})
}

//example: film/?format=csv&genre=fantasy,action,drama&rate=7-8.6&sort=minutes,rate,year&limit=9&offset=0
//
//example: SELECT f.id, f.name, f.genre, d.name, f.rate, f.year, f.minutes FROM film f JOIN director d
//ON f.director_id = d.id WHERE genre IN ('fantasy','action','drama') AND (rate >= 7 AND rate <= 8.6)
//ORDER BY minutes, rate, year LIMIT 5 OFFSET 0;

func (h *Handler) GetAllFilms(c *gin.Context) {
	var CSV []byte
	var params entities.QueryParams

	params.Format = c.Query("format")
	params.Genre = c.Query("genre")
	params.Rate = c.Query("rate")
	params.Sort = c.Query("sort")
	params.Limit = c.Query("limit")
	params.Offset = c.Query("offset")

	userId, _ := c.Get(userCtx)

	SQL, err := h.usecases.MakeQuery(params)
	if err != nil {
		logrus.Warnf("Attempt to get film list: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	filmList, err := h.usecases.GetFilmList(SQL)
	if err != nil {
		logrus.Warnf("Attempt to get film list: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

		return
	}

	switch params.Format {
	case "json":
		logrus.Infof("List is JSON successfully sended. User ID: %+v", userId)
		c.JSON(http.StatusOK, entities.GetAllFilmsResponse{
			Data: filmList,
		})
	case "csv":
		CSV, err = gocsv.MarshalBytes(filmList)
		if err != nil {
			logrus.Errorf("Attempt to get film list: %v. User ID: %v", err, userId)
			newResponse(c, http.StatusInternalServerError, Response{Message: MsgInternalServerErr, Details: err.Error()})

			return
		}
		logrus.Infof("List is CSV successfully sended. User ID: %+v", userId)
		c.Data(http.StatusOK, "csv", CSV)

	default:
		logrus.Warnf("Attempt to get film list: %v. User ID: %v", MsgProblemFormatOutputData, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgProblemFormatOutputData, Details: "Format is empty or incorrect"})

		return
	}
}

func (h *Handler) GetFilmByID(c *gin.Context) {
	userId, _ := c.Get(userCtx)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Warnf("Attempt to get film by ID: %v. User ID: %v", MsgInvalidIDParam, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgInvalidIDParam})
		return
	}

	film, err := h.usecases.GetFilmById(id)
	if err != nil {
		logrus.Warnf("Attempt to get film by ID: %v. User ID: %v", err, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})
		return
	}

	logrus.Infof("Movie successfully received. User ID: %v. Film ID: %v", userId, id)
	c.JSON(http.StatusOK, film)
}

func (h *Handler) AddToList(c *gin.Context) {
	var params entities.ListsComponents

	params.List = c.Query("list")
	params.Movie = c.Query("movieName")

	userId, _ := c.Get(userCtx)

	switch params.List {
	case "favouritelist":
		id, err := h.usecases.AddFilmToList(userId, params.Movie, repository.FavouriteTable)
		if err != nil {
			logrus.Warnf("Attempt to add movie to favourite list: %v. User ID: %v. Film: %v", err, userId, params.Movie)
			newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

			return
		}

		logrus.Infof("Movie successfully added to favourite list. User ID: %v. Film: %v. ID in list: %v", userId, params.Movie, id)
		c.JSON(http.StatusOK, map[string]interface{}{
			"film's id in my favourite list": id,
		})
	case "wishlist":
		id, err := h.usecases.AddFilmToList(userId, params.Movie, repository.WishlistTable)
		if err != nil {
			logrus.Warnf("Attempt to add movie to wish list: %v. User ID: %v. Film: %v", err, userId, params.Movie)
			newResponse(c, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()})

			return
		}

		logrus.Infof("Movie successfully added to wish list. User ID: %v. Film: %v. ID in list: %v", userId, params.Movie, id)
		c.JSON(http.StatusOK, map[string]interface{}{
			"film's id in my wish list": id,
		})

	default:
		logrus.Warnf("Attempt to add film to list: %v. User ID: %v", MsgProblemFieldList, userId)
		newResponse(c, http.StatusBadRequest, Response{Message: MsgProblemFieldList, Details: "List field is empty or incorrect"})

		return
	}
}
