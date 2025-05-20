package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ridebooking/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocationService struct {
	mock.Mock
}

func (m *MockLocationService) UpdateDriverLocation(ctx context.Context, req model.DriverLocationRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockLocationService) UpdateDriverAvailability(ctx context.Context, req model.DriverLocationRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockLocationService) GetAllAvailableDrivers(ctx context.Context) ([]model.DriverLocation, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.DriverLocation), args.Error(1)
}

func (m *MockLocationService) GetAllNearByDrivers(ctx context.Context, location model.Gelocation) ([]model.DriverLocation, error) {
	args := m.Called(ctx, location)
	return args.Get(0).([]model.DriverLocation), args.Error(1)
}

func TestUpdateDriverLocationHandler_Success(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	reqBody := model.DriverLocationRequest{
		DriverId: "D1",
		Location: model.Gelocation{X: 1.0, Y: 2.0},
	}
	mockService.On("UpdateDriverLocation", mock.Anything, reqBody).Return(nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/api/ridebooking/driver/location", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.UpdateDriverLocationHandler(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateDriverAvailabilityHandler_Success(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	reqBody := model.DriverLocationRequest{
		DriverId:    "D1",
		IsAvailable: true,
	}
	mockService.On("UpdateDriverAvailability", mock.Anything, reqBody).Return(nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/api/ridebooking/driver/availability", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.UpdateDriverAvailabilityHandler(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetAllAvailableDriversHandler_Success(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	mockDrivers := []model.DriverLocation{
		{DriverId: "D1", IsAvailable: true, Location: model.Gelocation{X: 1.0, Y: 2.0}},
	}
	mockService.On("GetAllAvailableDrivers", mock.Anything).Return(mockDrivers, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/driver/available", nil)
	rr := httptest.NewRecorder()

	handler.GetAllAvailableDriversHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetNearbyDriversHandler_Success(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	location := model.Gelocation{X: 10, Y: 20}
	mockDrivers := []model.DriverLocation{
		{DriverId: "D1", Location: location, IsAvailable: true},
	}
	mockService.On("GetAllNearByDrivers", mock.Anything, location).Return(mockDrivers, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/driver/nearby?x=10&y=20", nil)
	rr := httptest.NewRecorder()

	handler.GetNearbyDriversHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetNearbyDriversHandler_InvalidX(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/driver/nearby?x=abc&y=20", nil)
	rr := httptest.NewRecorder()

	handler.GetNearbyDriversHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetNearbyDriversHandler_InvalidY(t *testing.T) {
	mockService := new(MockLocationService)
	handler := NewLocationHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/driver/nearby?x=10&y=abc", nil)
	rr := httptest.NewRecorder()

	handler.GetNearbyDriversHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
