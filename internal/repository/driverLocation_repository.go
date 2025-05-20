package repository

import (
	"context"
	"errors"
	"fmt"
	"ridebooking/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverLocationRepository struct {
	Collection *mongo.Collection
}

func NewDriverLocationRepository(collection *mongo.Collection) *DriverLocationRepository {
	return &DriverLocationRepository{
		Collection: collection,
	}
}

func (locationRepo *DriverLocationRepository) CreateDriverLocation(ctx context.Context, driverLocation *model.DriverLocation) error {
	_, err := locationRepo.Collection.InsertOne(ctx, driverLocation)
	return err
}

func (locationRepo *DriverLocationRepository) UpdateDriverLocation(ctx context.Context, driverLocation *model.DriverLocation) error {
	var fetchLocation model.DriverLocation
	err := locationRepo.Collection.FindOne(ctx, bson.M{"driver_id": driverLocation.DriverId}).Decode(&fetchLocation)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if createErr := locationRepo.CreateDriverLocation(ctx, driverLocation); createErr != nil {
				return fmt.Errorf("failed to create driver location: %w", createErr)
			}
			return nil
		}
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"location":     driverLocation.Location,
			"last_updated": time.Now(),
		},
	}
	_, error := locationRepo.Collection.UpdateByID(ctx, fetchLocation.ID, update)
	if error != nil {
		return error
	}
	return nil
}

func (locationRepo *DriverLocationRepository) UpdateDriverAvailability(ctx context.Context, driverLocation *model.DriverLocation) error {
	var fetchLocation model.DriverLocation
	err := locationRepo.Collection.FindOne(ctx, bson.M{"driver_id": driverLocation.DriverId}).Decode(&fetchLocation)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"is_available": driverLocation.IsAvailable,
			"last_updated": time.Now(),
		},
	}
	_, error := locationRepo.Collection.UpdateByID(ctx, fetchLocation.ID, update)
	if error != nil {
		return error
	}
	return nil
}

func (locationRepo *DriverLocationRepository) GetAllAvailableDrivers(ctx context.Context) ([]model.DriverLocation, error) {
	cursor, err := locationRepo.Collection.Find(ctx, bson.M{"is_available": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var driverLocations []model.DriverLocation
	for cursor.Next(ctx) {
		var driverLocation model.DriverLocation
		if err := cursor.Decode(&driverLocation); err != nil {
			return nil, err
		}
		driverLocations = append(driverLocations, driverLocation)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return driverLocations, nil
}

func (locationRepo *DriverLocationRepository) RemoveDriverLocationById(ctx context.Context, driverId string) error {
	var fetchLocation model.DriverLocation
	err := locationRepo.Collection.FindOne(ctx, bson.M{"driver_id": driverId}).Decode(&fetchLocation)
	if err != nil {
		return err
	}
	_, error := locationRepo.Collection.DeleteOne(ctx, bson.M{"_id": fetchLocation.ID})
	if error != nil {
		return error
	}
	return nil
}
