package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DriverLocation struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	DriverId    string             `bson:"driver_id,omitempty"`
	IsAvailable bool               `bson:"is_available,omitempty"`
	Location    Gelocation         `bson:"location,omitempty"`
	LastUpdated time.Time          `bson:"last_updated,omitempty"`
}

type DriverLocationRequest struct {
	DriverId    string     `bson:"driver_id"`
	IsAvailable bool       `bson:"is_available"`
	Location    Gelocation `bson:"location"`
}
