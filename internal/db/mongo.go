package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client                   *mongo.Client
	UserCollection           *mongo.Collection
	TripCollection           *mongo.Collection
	DriverLocationCollection *mongo.Collection
)

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	Client = client
	UserCollection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("USER_COLLECTION"))
	TripCollection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("TRIP_COLLECTION"))
	DriverLocationCollection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("DRIVER_LOCATION_COLLECTION"))
}
