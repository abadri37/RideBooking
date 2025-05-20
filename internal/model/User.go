package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType string

const (
	Rider  UserType = "Rider"
	Driver UserType = "Driver"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"user_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	FirstName string             `bson:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty"`
	Type      UserType           `bson:"type,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Gelocation struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
}

type UserRequest struct {
	UserId    string   `json:"user_id" bson:"user_id"`
	Email     string   `json:"email"    bson:"email"`
	Password  string   `json:"password" bson:"password"`
	FirstName string   `json:"first_name" bson:"first_name"`
	LastName  string   `json:"last_name"  bson:"last_name"`
	Type      UserType `json:"type"       bson:"type"`
}

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	Token string `bson:"token"`
}
