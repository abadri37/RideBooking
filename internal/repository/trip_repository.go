package repository

import (
	"context"
	"ridebooking/internal/model"
	"ridebooking/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TripRepository struct {
	Collection *mongo.Collection
}

func NewTripRepository(collection *mongo.Collection) *TripRepository {
	return &TripRepository{
		Collection: collection,
	}
}

func (tripRepo *TripRepository) CreateTrip(ctx context.Context, trip *model.Trip) error {
	_, err := tripRepo.Collection.InsertOne(ctx, trip)
	return err
}

func (tripRepo *TripRepository) UpdateTrip(ctx context.Context, trip *model.Trip) error {
	var fetchTrip model.Trip
	err := tripRepo.Collection.FindOne(ctx, bson.M{"trip_id": trip.TripId}).Decode(&fetchTrip)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"driver_id":      utils.GetValueOrFallback(trip.DriverId, fetchTrip.DriverId),
			"rider_id":       utils.GetValueOrFallback(trip.RiderId, fetchTrip.RiderId),
			"start_location": utils.GetLocationOrFallback(trip.StartLocation, fetchTrip.StartLocation),
			"end_location":   utils.GetLocationOrFallback(trip.EndLocation, fetchTrip.EndLocation),
			"total_distance": utils.GetFloatOrFallback(trip.TotalDistance, fetchTrip.TotalDistance),
			"status":         utils.GetValueOrFallbackTrip(trip.Status, fetchTrip.Status),
			"start_time":     utils.GetTimeOrFallback(trip.StartTime, fetchTrip.StartTime),
			"end_time":       utils.GetTimeOrFallback(trip.EndTime, fetchTrip.EndTime),
		},
	}
	_, error := tripRepo.Collection.UpdateByID(ctx, fetchTrip.ID, update)
	if error != nil {
		return error
	}
	return nil
}

func (tripRepo *TripRepository) FetchTripByDriverId(ctx context.Context, driverId string) ([]model.Trip, error) {
	cursor, err := tripRepo.Collection.Find(ctx, bson.M{"driver_id": driverId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var trips []model.Trip
	for cursor.Next(ctx) {
		var trip model.Trip
		if err := cursor.Decode(&trip); err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return trips, nil
}

func (tripRepo *TripRepository) FetchTripByRiderId(ctx context.Context, riderId string) ([]model.Trip, error) {
	cursor, err := tripRepo.Collection.Find(ctx, bson.M{"rider_id": riderId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var trips []model.Trip
	for cursor.Next(ctx) {
		var trip model.Trip
		if err := cursor.Decode(&trip); err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return trips, nil
}

func (tripRepo *TripRepository) FetchTripByTripId(ctx context.Context, tripId string) (*model.Trip, error) {
	var trip model.Trip
	err := tripRepo.Collection.FindOne(ctx, bson.M{"trip_id": tripId}).Decode(&trip)
	if err != nil {
		return nil, err
	}
	return &trip, nil
}
