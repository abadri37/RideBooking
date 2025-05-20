package utils

import (
	"log"
	"math"
	"ridebooking/internal/model"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pasword string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedBytes)
}
func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func GetUniqueId() string {
	id := uuid.New()
	return id.String()
}

func CalculateDistance(a, b model.Gelocation) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func GetValueOrFallback(value, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
}

func GetValueOrFallbackTrip(value, fallback model.TripStatus) model.TripStatus {
	if value != "" {
		return value
	}
	return fallback
}

// For location fields (assuming Gelocation struct)
func GetLocationOrFallback(value, fallback model.Gelocation) model.Gelocation {
	if (value != model.Gelocation{}) {
		return value
	}
	return fallback
}

// For float64 fields (e.g., total_distance)
func GetFloatOrFallback(value, fallback float64) float64 {
	if value != 0 {
		return value
	}
	return fallback
}

// For time.Time fields (e.g., start_time, end_time)
func GetTimeOrFallback(value, fallback time.Time) time.Time {
	if !value.IsZero() {
		return value
	}
	return fallback
}
