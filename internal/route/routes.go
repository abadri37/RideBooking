package route

import (
	"ridebooking/internal/handler"
	"ridebooking/internal/middleware"

	_ "ridebooking/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(userHandler *handler.UserHandler, tripHandler *handler.TripHandler, locationHandler *handler.LocationHandler) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) // Serve Swagger UI

	// Register the routes for the user service
	r.HandleFunc("/login", userHandler.UserLoginHandler).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUserHandler).Methods("POST")

	// Protected routes (JWT Middleware)
	protected := r.PathPrefix("/api/ridebooking/").Subrouter()
	protected.Use(middleware.JwtMiddleWare)

	protected.HandleFunc("/user/emailId", userHandler.GetUserByEmailHandler).Methods("GET")
	protected.HandleFunc("/user/id", userHandler.GetUserByIdHandler).Methods("GET")
	protected.HandleFunc("/user", userHandler.UpdateUserHandler).Methods("PUT")
	protected.HandleFunc("/user/emailId", userHandler.RemoveUserByEmailHandler).Methods("DELETE")

	protected.HandleFunc("/driver/location", locationHandler.UpdateDriverLocationHandler).Methods("PUT")
	protected.HandleFunc("/driver/availability", locationHandler.UpdateDriverAvailabilityHandler).Methods("PUT")
	protected.HandleFunc("/driver/available", locationHandler.GetAllAvailableDriversHandler).Methods("GET")
	protected.HandleFunc("/driver/nearby", locationHandler.GetNearbyDriversHandler).Methods("GET")

	protected.HandleFunc("/trip", tripHandler.CreateTripHandler).Methods("POST")
	protected.HandleFunc("/trip", tripHandler.UpdateTripHandler).Methods("PUT")
	protected.HandleFunc("/trip", tripHandler.GetTripByIdHandler).Methods("GET")
	protected.HandleFunc("/trip/driver", tripHandler.GetTripByDriverIdHandler).Methods("GET")
	protected.HandleFunc("/trip/rider", tripHandler.GetTripByRiderIdHandler).Methods("GET")

	return r
}
