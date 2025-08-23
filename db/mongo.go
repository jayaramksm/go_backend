package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	// MongoDB connection URI
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jayaram:jayaram@cluster0.vlqwqiz.mongodb.net/go_crud?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("go_crud")
	log.Println("✅ Connected to MongoDB")
}
