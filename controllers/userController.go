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

// Create User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
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
	user.UserID = uidStr
	log.Println("user_id==>", userId)
	user.ID = primitive.NewObjectID()

	collection := db.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	} else {
		log.Println("Insert result:", res.InsertedID)
	}

	// c.JSON(http.StatusCreated, user)
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"debug": "User successfully inserted into DB",
	})
}

// Get All Users
func GetUsers(c *gin.Context) {
	collection := db.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userId, _ := c.Get("user_id")
	uidStr := userId.(string)
	cursor, err := collection.Find(ctx, bson.M{"_user_id": uidStr})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(ctx)
	log.Println("cursor===>_", cursor)
	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Get Single User
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := db.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// and got logged-in user_id from middleware
	uidStr, _ := c.Get("user_id")     // comes from JWT middleware
	loggedInUserID := uidStr.(string) // convert interface{} → string

	// Compare DB user.UserID with logged-in user ID
	if user.UserID != loggedInUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update User
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("users")

	var dbuser models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&dbuser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	uidStr, _ := c.Get("user_id")     // comes from JWT middleware
	loggedInUserID := uidStr.(string) // convert interface{} → string

	// Compare DB user.UserID with logged-in user ID
	if dbuser.UserID != loggedInUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// Delete User
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := db.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dbuser models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&dbuser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	uidStr, _ := c.Get("user_id")     // comes from JWT middleware
	loggedInUserID := uidStr.(string) // convert interface{} → string

	// Compare DB user.UserID with logged-in user ID
	if dbuser.UserID != loggedInUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
