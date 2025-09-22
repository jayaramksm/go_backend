package main

import (
	"go_backend/db"
	"go_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	db.Connect()

	r := gin.Default()

	// // Routes
	// r.POST("/users", handlers.CreateUser)
	// r.GET("/users", handlers.GetUsers)
	// r.GET("/users/:id", handlers.GetUser)
	// r.PUT("/users/:id", handlers.UpdateUser)
	// r.DELETE("/users/:id", handlers.DeleteUser)

	// Register routes
	routes.UserRoutes(r)
	routes.AuthRoutes(r)

	r.Run(":8080") // Start server on localhost:8080
}
