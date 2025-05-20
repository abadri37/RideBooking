package handler

import (
	"encoding/json"
	"net/http"
	_ "ridebooking/docs"
	"ridebooking/internal/model"
	"ridebooking/internal/service"
	"strconv"
)

type LocationHandler struct {
	locationService service.LocationService
}

func NewLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService: locationService,
	}
}

// @Summary Update driver location
// @Description Updates the location of a driver
// @Tags drivers
// @Accept json
// @Produce json
// @Param location body model.DriverLocationRequest true "Driver location details"
// @Success 202 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/driver/location [put]
func (h *LocationHandler) UpdateDriverLocationHandler(w http.ResponseWriter, r *http.Request) {
	var driverLocationRequest model.DriverLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&driverLocationRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.locationService.UpdateDriverLocation(r.Context(), driverLocationRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Driver Location Updated successfully"})
}

// @Summary Update driver availability
// @Description Updates the availability status of a driver
// @Tags drivers
// @Accept json
// @Produce json
// @Param availability body model.DriverLocationRequest true "Driver availability details"
// @Success 202 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/driver/availability [put]
func (h *LocationHandler) UpdateDriverAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	var driverLocationRequest model.DriverLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&driverLocationRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.locationService.UpdateDriverAvailability(r.Context(), driverLocationRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Driver Availability Updated successfully"})
}

// @Summary Get all available drivers
// @Description Retrieves a list of all drivers who are currently available
// @Tags drivers
// @Produce json
// @Success 200 {array} model.DriverLocation
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/driver/available [get]
func (h *LocationHandler) GetAllAvailableDriversHandler(w http.ResponseWriter, r *http.Request) {

	drivers, err := h.locationService.GetAllAvailableDrivers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drivers)
}

// @Summary Get nearby drivers
// @Description Finds available drivers near a specific location
// @Tags drivers
// @Produce json
// @Param x query number true "X coordinate"
// @Param y query number true "Y coordinate"
// @Success 200 {array} model.DriverLocation
// @Failure 400 {string} string "Invalid coordinates"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/driver/nearby [get]
func (h *LocationHandler) GetNearbyDriversHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	xStr := query.Get("x")
	yStr := query.Get("y")

	x, err := strconv.ParseFloat(xStr, 64)
	if err != nil {
		http.Error(w, "Invalid x value", http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseFloat(yStr, 64)
	if err != nil {
		http.Error(w, "Invalid y value", http.StatusBadRequest)
		return
	}
	location := model.Gelocation{X: x, Y: y}

	drivers, err := h.locationService.GetAllNearByDrivers(r.Context(), location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drivers)
}
