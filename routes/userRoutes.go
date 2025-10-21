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
	// {
	user.GET("/", middleware.RoleMiddleware("admin", "user"), controllers.GetUsers)
	user.POST("/", middleware.RoleMiddleware("admin", "user"), controllers.CreateUser)
	user.GET("/:id", middleware.RoleMiddleware("admin", "user"), controllers.GetUserByID)
	user.PUT("/:id", middleware.RoleMiddleware("admin", "user"), controllers.UpdateUser)
	user.DELETE("/:id", middleware.RoleMiddleware("admin"), controllers.DeleteUser)
	// }
}
