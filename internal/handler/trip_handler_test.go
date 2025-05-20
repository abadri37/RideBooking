package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"ridebooking/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTripService mocks the service.TripService interface
type MockTripService struct {
	mock.Mock
}

func (m *MockTripService) CreateTrip(ctx context.Context, req model.TripRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

func (m *MockTripService) UpdateTrip(ctx context.Context, req model.TripRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockTripService) GetTripById(ctx context.Context, id string) (*model.Trip, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (m *MockTripService) GetDriverTrips(ctx context.Context, driverId string) ([]model.Trip, error) {
	args := m.Called(ctx, driverId)
	return args.Get(0).([]model.Trip), args.Error(1)
}

func (m *MockTripService) GetRiderTrips(ctx context.Context, riderId string) ([]model.Trip, error) {
	args := m.Called(ctx, riderId)
	return args.Get(0).([]model.Trip), args.Error(1)
}

func TestCreateTripHandler_Success(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	reqBody := model.TripRequest{RiderId: "r123", DriverId: "d456"}
	jsonBody, _ := json.Marshal(reqBody)

	mockService.On("CreateTrip", mock.Anything, reqBody).Return("trip123", nil)

	req := httptest.NewRequest(http.MethodPost, "/api/ridebooking/trip", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	handler.CreateTripHandler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Trip Created successfully trip123")
}

func TestGetTripByIdHandler_Success(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	trip := &model.Trip{
		TripId:    "trip123",
		RiderId:   "r123",
		DriverId:  "d456",
		Status:    model.TripOngoing,
		StartTime: time.Now(),
	}
	mockService.On("GetTripById", mock.Anything, "trip123").Return(trip, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/trip?tripId=trip123", nil)
	w := httptest.NewRecorder()

	handler.GetTripByIdHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "trip123")
}

func TestUpdateTripHandler_Success(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	reqBody := model.TripRequest{
		TripId:   "trip123",
		RiderId:  "r123",
		DriverId: "d456",
	}
	jsonBody, _ := json.Marshal(reqBody)

	mockService.On("UpdateTrip", mock.Anything, reqBody).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/ridebooking/trip", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	handler.UpdateTripHandler(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Contains(t, w.Body.String(), "Trip Updated successfully trip123")
}

func TestGetTripByDriverIdHandler_Success(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	driverId := "d456"
	tripList := []model.Trip{
		{TripId: "t1", DriverId: driverId},
		{TripId: "t2", DriverId: driverId},
	}

	mockService.On("GetDriverTrips", mock.Anything, driverId).Return(tripList, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/trip/driver?driverId=d456", nil)
	w := httptest.NewRecorder()

	handler.GetTripByDriverIdHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "t1")
	assert.Contains(t, w.Body.String(), "t2")
}

func TestGetTripByRiderIdHandler_Success(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	riderId := "r123"
	tripList := []model.Trip{
		{TripId: "t3", RiderId: riderId},
		{TripId: "t4", RiderId: riderId},
	}

	mockService.On("GetRiderTrips", mock.Anything, riderId).Return(tripList, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/trip/rider?riderId=r123", nil)
	w := httptest.NewRecorder()

	handler.GetTripByRiderIdHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "t3")
	assert.Contains(t, w.Body.String(), "t4")
}

func TestCreateTripHandler_Failure(t *testing.T) {
	mockService := new(MockTripService)
	handler := NewTripHandler(mockService)

	reqBody := model.TripRequest{RiderId: "r123", DriverId: "d456"}
	jsonBody, _ := json.Marshal(reqBody)

	mockService.On("CreateTrip", mock.Anything, reqBody).Return("", errors.New("failed to create trip"))

	req := httptest.NewRequest(http.MethodPost, "/api/ridebooking/trip", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	handler.CreateTripHandler(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to create trip")
}
