package services

import (
	"math/rand"
	"time"

	"github.com/jc80800/badminton-chaos/internal/models"
)

// QuoteService handles quote-related business logic
type QuoteService struct {
	quotes []string
	rng    *rand.Rand
}

func NewQuoteService() *QuoteService {
	// Create a new random number generator with current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &QuoteService{
		quotes: models.GetRandomQuotes(),
		rng:    rng,
	}
}

func (s *QuoteService) GetRandomQuote() string {
	return s.quotes[s.rng.Intn(len(s.quotes))]
}
