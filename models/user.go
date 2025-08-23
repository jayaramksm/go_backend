package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID string             `json:"user_id,omitempty" bson:"_user_id,omitempty"`
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
	Age    int                `json:"age" bson:"age"`
}
