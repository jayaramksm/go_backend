package main

import (
	"go_backend/db"
	"go_backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Connect to MongoDB
	err := godotenv.Load()
	db.Connect()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	port := os.Getenv("PORT")
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

	r.Run(":" + port) // Start server on localhost:8080
}
