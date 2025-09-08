package main

import (
	"log"

	"github.com/jc80800/badminton-chaos/internal/handlers"
	"github.com/jc80800/badminton-chaos/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database
	dbService, err := services.NewDatabaseService()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbService.Close()

	// Initialize services
	excuseService := services.NewExcuseService(dbService.GetPool())

	// Create Echo instance
	e := echo.New()

	// Add CORS middleware
	e.Use(middleware.CORS())

	// Initialize handlers
	helloHandler := handlers.NewHelloHandler()
	excuseHandler := handlers.NewExcuseHandler(excuseService)

	// Routes
	e.GET("/", helloHandler.Hello)
	e.GET("/random", excuseHandler.GetRandomExcuse)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
