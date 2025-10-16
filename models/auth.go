package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email" binding:"required"`
	Password string             `bson:"password" json:"password" binding:"required"`
	Role     string             `bson:"role" json:"role" binding:"required"`
}
