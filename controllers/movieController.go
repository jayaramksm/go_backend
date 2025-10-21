package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"go_backend/db"
	"go_backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMovies(c *gin.Context) {
	collection := db.DB.Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userId, _ := c.Get("user_id")
	uidStr := userId.(string)
	cursor, err := collection.Find(ctx, bson.M{"_user_id": uidStr})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	defer cursor.Close(ctx)
	var movies []models.Movies
	if err = cursor.All(ctx, &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse users"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func CreateMovies(c *gin.Context) {
	var movie models.Movies
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Type assert to string
	uidStr := userId.(string)
	movie.UserID = uidStr
	log.Println("user_id==>", userId)
	movie.ID = primitive.NewObjectID()

	collection := db.DB.Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	} else {
		log.Println("Insert result:", res.InsertedID)
	}

	// c.JSON(http.StatusCreated, user)
	c.JSON(http.StatusCreated, gin.H{
		"user":  movie,
		"debug": "movie successfully inserted into DB",
	})
}
