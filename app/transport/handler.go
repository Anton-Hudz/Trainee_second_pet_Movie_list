package transport

import (
	"github.com/gin-contrib/cors"
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

// func (h *Handler) addCorsHeaders(router *gin.Engine) {
// 	// Добавляем middleware для установки заголовка Access-Control-Allow-Origin
// 	router.Use(func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	})
// }

// func (h *Handler) InitRoutes() *gin.Engine {
// 	router := gin.New()
// 	users := router.Group("/user")
// 	{
// 		users.POST("/sign-up", h.CreateUser)
// 		users.POST("/sign-in", h.LogIn)
// 		users.DELETE("/", h.Logout)
// 	}
// 	film := router.Group("/film", h.UserIdentity)
// 	{
// 		film.POST("/", h.CreateFilm)
// 		film.GET("/", h.GetAllFilms)
// 		film.GET("/:id", h.GetFilmByID)
// 		film.POST("/:id", h.AddToList)
// 	}

// 	return router
// }

//FOR AXIOS&CORS
func (h *Handler) InitRoutes() *gin.Engine {
	// Создаем конфигурацию CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	corsConfig.AllowMethods = []string{"POST", "PUT", "DELETE"}

	// // Создаем новый обработчик CORS с указанной конфигурацией
	c := cors.New(corsConfig)

	// Создаем новый маршрутизатор Gin
	router := gin.Default()

	// ---------------
	// router.OPTIONS("/*cors", func(c *gin.Context) {
	// 	c.Status(200)
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	// 	c.Header("Content-Type", "application/json")
	// 	c.AbortWithStatus(200)
	// })
	// ----------

	// Используем обработчик CORS в качестве middleware
	// router.Use(CorsMiddleware())

	// --------------------------
	router.Use(c)
	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
	// 	c.Next()
	// })

	// ---------------------

	// Определяем маршруты
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
