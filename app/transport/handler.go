package transport

import (
	"github.com/Anton-Hudz/MovieList/app/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.UseCase
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// userId, _ := c.Get(userCtx)
	users := router.Group("/user")
	{
		users.POST("/sign-up", h.CreateUser)
		users.POST("/sign-in", h.LogIn)
		users.DELETE("/", h.LogOut)
	}
	film := router.Group("/film", h.UserIdentity)
	{
		film.POST("/", h.CreateFilm)
		film.GET("/", h.GetAllFilms)
		film.GET("/:id", h.GetFilmByID)
		film.POST("/:id", h.AddToList)
	}

	return router
}
