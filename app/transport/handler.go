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
	users := router.Group("/users")
	{
		users.POST("/sign-up", h.CreateUser)
		users.POST("/sign-in", h.LogIn)
		users.DELETE("/", h.LogOut)
	}
	film := router.Group("/film")
	{
		film.POST("/", h.CreateFilm)
		film.GET("/", h.GetAllFilms)
		film.GET("/:id", h.GetFilmByID)
	}
	router.POST("/favourites", h.AddToFavourite)
	router.POST("/wishlist", h.AddToWishlist)
	router.GET("/get-csv-file", h.GetCSVFile)

	return router
}
