package service

import (
	"context"
	"ridebooking/internal/model"
	"ridebooking/internal/repository"
	"ridebooking/internal/utils"
	"time"
)

type LocationService interface {
	UpdateDriverLocation(ctx context.Context, driverLocationReq model.DriverLocationRequest) error
	UpdateDriverAvailability(ctx context.Context, driverLocationReq model.DriverLocationRequest) error
	GetAllAvailableDrivers(ctx context.Context) ([]model.DriverLocation, error)
	GetAllNearByDrivers(ctx context.Context, location model.Gelocation) ([]model.DriverLocation, error)
}

type LocationServiceImpl struct {
	driverLocationRepo *repository.DriverLocationRepository
}

func NewLocationServiceImpl(repo *repository.DriverLocationRepository) *LocationServiceImpl {
	return &LocationServiceImpl{
		driverLocationRepo: repo,
	}
}

func (LocationService *LocationServiceImpl) UpdateDriverLocation(ctx context.Context, driverLocationReq model.DriverLocationRequest) error {

	driverLocation := &model.DriverLocation{
		DriverId:    driverLocationReq.DriverId,
		IsAvailable: driverLocationReq.IsAvailable,
		Location:    driverLocationReq.Location,
		LastUpdated: time.Now(),
	}
	err := LocationService.driverLocationRepo.UpdateDriverLocation(ctx, driverLocation)
	if err != nil {
		return err
	}
	return nil
}

func (LocationService *LocationServiceImpl) UpdateDriverAvailability(ctx context.Context, driverLocationReq model.DriverLocationRequest) error {

	driverLocation := &model.DriverLocation{
		DriverId:    driverLocationReq.DriverId,
		IsAvailable: driverLocationReq.IsAvailable,
		LastUpdated: time.Now(),
	}
	err := LocationService.driverLocationRepo.UpdateDriverAvailability(ctx, driverLocation)
	if err != nil {
		return err
	}
	return nil
}

func (LocationService *LocationServiceImpl) GetAllAvailableDrivers(ctx context.Context) ([]model.DriverLocation, error) {
	driverLocations, err := LocationService.driverLocationRepo.GetAllAvailableDrivers(ctx)
	if err != nil {
		return nil, err
	}
	return driverLocations, nil
}

func (LocationService *LocationServiceImpl) GetAllNearByDrivers(ctx context.Context, location model.Gelocation) ([]model.DriverLocation, error) {
	driverLocations, err := LocationService.driverLocationRepo.GetAllAvailableDrivers(ctx)
	if err != nil {
		return nil, err
	}
	var nearByDrivers []model.DriverLocation
	for _, drivers := range driverLocations {
		dist := utils.CalculateDistance(location, drivers.Location)
		if dist <= 5 { /**validating max radius of 5 KM*/
			nearByDrivers = append(nearByDrivers, drivers)
		}
	}
	return nearByDrivers, nil
}
