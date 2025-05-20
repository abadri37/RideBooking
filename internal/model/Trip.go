package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripStatus string

const (
	TripPending   TripStatus = "pending"
	TripAccepted  TripStatus = "accepted"
	TripOngoing   TripStatus = "ongoing"
	TripCancelled TripStatus = "cancelled"
	TripCompleted TripStatus = "completed"
)

type Trip struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	TripId        string             `bson:"trip_id,omitempty"`
	DriverId      string             `bson:"driver_id,omitempty"`
	RiderId       string             `bson:"rider_id,omitempty"`
	StartLocation Gelocation         `bson:"start_location,omitempty"`
	EndLocation   Gelocation         `bson:"end_location,omitempty"`
	TotalDistance float64            `bson:"total_distance,omitempty"`
	Status        TripStatus         `bson:"status,omitempty"`
	StartTime     time.Time          `bson:"start_time,omitempty"`
	EndTime       time.Time          `bson:"end_time,omitempty"`
}

type TripRequest struct {
	TripId        string     `bson:"trip_id"`
	DriverId      string     `bson:"driver_id"`
	RiderId       string     `bson:"rider_id"`
	StartLocation Gelocation `bson:"start_location"`
	EndLocation   Gelocation `bson:"end_location"`
	TotalDistance float64    `bson:"total_distance"`
	Status        TripStatus `bson:"status"`
	StartTime     time.Time  `bson:"start_time"`
	EndTime       time.Time  `bson:"end_time"`
}
