package routes

import (
	"go_backend/controllers"
	"go_backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// Grouping routes under /users
	user := r.Group("/api/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}
