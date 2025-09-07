package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("/health", health)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Badminton Chaos!")
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "healthy",
		"message": "Badminton Chaos API is running",
	})
}
