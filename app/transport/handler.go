package transport

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserUseCase
	FilmUseCase
}

func NewHandler(uu UserUseCase, fu FilmUseCase) *Handler {
	return &Handler{
		UserUseCase: uu,
		FilmUseCase: fu,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	router.Use(CORSMiddleware())

	users := router.Group("/user")
	{
		users.POST("/sign-up", h.CreateUser)
		users.POST("/sign-in", h.LogIn)
		users.DELETE("/", h.Logout)
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
