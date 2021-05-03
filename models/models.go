package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var err error
var ctx context.Context

var moviesCollection *mongo.Collection
var userCollection *mongo.Collection
var reviewCollection *mongo.Collection

func ConnectAndInitialize() {
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@ratingapp.wlmbz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer Client.Disconnect(ctx)

	moviesCollection = Client.Database("movieRatingServer").Collection("movies")
	// defer Client.Disconnect(ctx)
	reviewCollection = Client.Database("movieRatingServer").Collection("reviews")

	userCollection = Client.Database("movieRatingServer").Collection("users")

}
