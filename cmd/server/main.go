package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ridebooking/internal/db"
	"ridebooking/internal/handler"
	"ridebooking/internal/repository"
	"ridebooking/internal/route"
	"ridebooking/internal/service"
	"syscall"
	"time"

	_ "ridebooking/docs"

	"github.com/joho/godotenv"
)

// @title Ride Booking API
// @version 1.0
// @description This is the user management API for ride booking service.
// @contact.name API Support
// @contact.email support@ridebooking.com
// @host localhost:8090
// @BasePath /
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Sample Program")
	fmt.Println("🚕 Starting Ride Booking Service...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using default environment variables.")
	}
	db.InitMongoDB()
	/*repositories **/
	userRepo := repository.NewUserRepository(db.UserCollection)
	locationRepo := repository.NewDriverLocationRepository(db.DriverLocationCollection)
	tripRepo := repository.NewTripRepository(db.TripCollection)

	/*services **/
	userService := service.NewUserService(userRepo, locationRepo)
	tripService := service.NewTripService(tripRepo)
	locationService := service.NewLocationServiceImpl(locationRepo)

	/*Hanlders*/
	userHandler := handler.NewUserHandler(userService)
	tripHandler := handler.NewTripHandler(tripService)
	locationHandler := handler.NewLocationHandler(locationService)

	// Registering routes
	router := route.RegisterRoutes(userHandler, tripHandler, locationHandler)

	// Start server
	server := &http.Server{
		Addr:    ":8090",
		Handler: router,
	}
	// Graceful shutdown handling
	go func() {
		log.Printf("🚀 Server is listening on http://localhost%s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Listen for interrupt signal (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("🛑 Shutting down server...")

	// Give the server a deadline to finish current operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited gracefully.")

}
