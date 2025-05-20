package service

import (
	"context"
	"ridebooking/internal/model"
	"ridebooking/internal/repository"
	"ridebooking/internal/utils"
)

type TripService interface {
	CreateTrip(ctx context.Context, TripRequest model.TripRequest) (string, error)
	UpdateTrip(ctx context.Context, TripRequest model.TripRequest) error
	GetRiderTrips(ctx context.Context, riderId string) ([]model.Trip, error)
	GetDriverTrips(ctx context.Context, driverId string) ([]model.Trip, error)
	GetTripById(ctx context.Context, tripId string) (*model.Trip, error)
}

type TripServiceImpl struct {
	tripRepo *repository.TripRepository
}

func NewTripService(repo *repository.TripRepository) *TripServiceImpl {
	return &TripServiceImpl{tripRepo: repo}
}

func (tripService *TripServiceImpl) CreateTrip(ctx context.Context, TripRequest model.TripRequest) (string, error) {
	trip := &model.Trip{
		TripId:        utils.GetUniqueId(),
		DriverId:      TripRequest.DriverId,
		RiderId:       TripRequest.RiderId,
		StartLocation: TripRequest.StartLocation,
		EndLocation:   TripRequest.EndLocation,
		TotalDistance: TripRequest.TotalDistance,
		Status:        model.TripPending,
	}
	err := tripService.tripRepo.CreateTrip(ctx, trip)
	if err != nil {
		return "", err
	}
	return trip.TripId, nil
}

func (tripService *TripServiceImpl) UpdateTrip(ctx context.Context, TripRequest model.TripRequest) error {
	trip := &model.Trip{
		TripId:        TripRequest.TripId,
		DriverId:      TripRequest.DriverId,
		RiderId:       TripRequest.RiderId,
		StartLocation: TripRequest.StartLocation,
		EndLocation:   TripRequest.EndLocation,
		TotalDistance: TripRequest.TotalDistance,
		Status:        TripRequest.Status,
	}
	err := tripService.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		return err
	}
	return nil
}

func (tripService *TripServiceImpl) GetRiderTrips(ctx context.Context, riderId string) ([]model.Trip, error) {
	trips, err := tripService.tripRepo.FetchTripByRiderId(ctx, riderId)
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (tripService *TripServiceImpl) GetDriverTrips(ctx context.Context, driverId string) ([]model.Trip, error) {
	trips, err := tripService.tripRepo.FetchTripByDriverId(ctx, driverId)
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (tripService *TripServiceImpl) GetTripById(ctx context.Context, tripId string) (*model.Trip, error) {
	trips, err := tripService.tripRepo.FetchTripByTripId(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return trips, nil
}
