package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler handles hello world requests
type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Badminton Chaos!")
}
