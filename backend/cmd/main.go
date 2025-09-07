package main

import (
	"github.com/jc80800/badminton-chaos/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Add CORS middleware
	e.Use(middleware.CORS())

	// Initialize handlers
	helloHandler := handlers.NewHelloHandler()
	quoteHandler := handlers.NewQuoteHandler()

	// Routes
	e.GET("/", helloHandler.Hello)
	e.GET("/random", quoteHandler.GetRandomQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
