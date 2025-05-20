package handler

import (
	"encoding/json"
	"net/http"
	_ "ridebooking/docs"
	"ridebooking/internal/model"
	"ridebooking/internal/service"
)

type TripHandler struct {
	tripService service.TripService
}

func NewTripHandler(tripService service.TripService) *TripHandler {
	return &TripHandler{
		tripService: tripService,
	}
}

// @Summary Create a new trip
// @Description Creates a trip with provided trip details
// @Tags trips
// @Accept json
// @Produce json
// @Param trip body model.TripRequest true "Trip data"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/trip [post]
func (h *TripHandler) CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	var tripRequest model.TripRequest
	if err := json.NewDecoder(r.Body).Decode(&tripRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	tripId, err := h.tripService.CreateTrip(r.Context(), tripRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Trip Created successfully " + tripId})
}

// @Summary Update a trip
// @Description Updates trip details
// @Tags trips
// @Accept json
// @Produce json
// @Param trip body model.TripRequest true "Updated trip data"
// @Success 202 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/trip [put]
func (h *TripHandler) UpdateTripHandler(w http.ResponseWriter, r *http.Request) {
	var tripRequest model.TripRequest
	if err := json.NewDecoder(r.Body).Decode(&tripRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.tripService.UpdateTrip(r.Context(), tripRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Trip Updated successfully " + tripRequest.TripId})
}

// @Summary Get trip by ID
// @Description Retrieves a trip by its unique trip ID
// @Tags trips
// @Produce json
// @Param tripId query string true "Trip ID"
// @Success 200 {object} model.Trip
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/trip [get]
func (h *TripHandler) GetTripByIdHandler(w http.ResponseWriter, r *http.Request) {

	tripId := r.URL.Query().Get("tripId")
	trip, err := h.tripService.GetTripById(r.Context(), tripId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trip)
}

// @Summary Get trips by driver ID
// @Description Retrieves all trips associated with a specific driver
// @Tags trips
// @Produce json
// @Param driverId query string true "Driver ID"
// @Success 200 {array} model.Trip
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/trip/driver [get]
func (h *TripHandler) GetTripByDriverIdHandler(w http.ResponseWriter, r *http.Request) {

	driverId := r.URL.Query().Get("driverId")
	trip, err := h.tripService.GetDriverTrips(r.Context(), driverId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trip)
}

// @Summary Get trips by rider ID
// @Description Retrieves all trips associated with a specific rider
// @Tags trips
// @Produce json
// @Param riderId query string true "Rider ID"
// @Success 200 {array} model.Trip
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/trip/rider [get]
func (h *TripHandler) GetTripByRiderIdHandler(w http.ResponseWriter, r *http.Request) {

	riderId := r.URL.Query().Get("riderId")
	trip, err := h.tripService.GetRiderTrips(r.Context(), riderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trip)
}
