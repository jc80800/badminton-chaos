package handlers

import (
	"net/http"

	"github.com/jc80800/badminton-chaos/internal/models"
	"github.com/jc80800/badminton-chaos/internal/services"
	"github.com/labstack/echo/v4"
)

// QuoteHandler handles quote-related HTTP requests
type QuoteHandler struct {
	quoteService *services.QuoteService
}

func NewQuoteHandler() *QuoteHandler {
	return &QuoteHandler{
		quoteService: services.NewQuoteService(),
	}
}

func (h *QuoteHandler) GetRandomQuote(c echo.Context) error {
	randomQuote := h.quoteService.GetRandomQuote()

	response := models.Response{
		Message: randomQuote,
	}

	return c.JSON(http.StatusOK, response)
}
