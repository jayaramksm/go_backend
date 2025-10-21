package routes

import (
	"go_backend/controllers"
	"go_backend/middleware"

	"github.com/gin-gonic/gin"
)

func MoviesRouters(r *gin.Engine) {

	movie := r.Group("/api/movies")
	movie.Use(middleware.AuthMiddleware())
	movie.GET("/", middleware.RoleMiddleware("admin", "user"), controllers.GetMovies)
	movie.POST("/", middleware.RoleMiddleware("admin", "user"), controllers.CreateMovies)
}
