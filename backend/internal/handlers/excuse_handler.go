package handlers

import (
	"net/http"

	"github.com/jc80800/badminton-chaos/internal/models"
	"github.com/jc80800/badminton-chaos/internal/services"
	"github.com/labstack/echo/v4"
)

// ExcuseHandler handles excuse-related HTTP requests
type ExcuseHandler struct {
	excuseService *services.ExcuseService
}

func NewExcuseHandler(excuseService *services.ExcuseService) *ExcuseHandler {
	return &ExcuseHandler{
		excuseService: excuseService,
	}
}

func (h *ExcuseHandler) GetRandomExcuse(c echo.Context) error {
	randomExcuse, err := h.excuseService.GetRandomExcuse()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to get random excuse",
		})
	}

	response := models.Response{
		Message: randomExcuse,
	}

	return c.JSON(http.StatusOK, response)
}
