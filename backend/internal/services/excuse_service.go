package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jc80800/badminton-chaos/internal/models"
)

// ExcuseService handles excuse-related business logic
type ExcuseService struct {
	db *pgxpool.Pool
}

func NewExcuseService(db *pgxpool.Pool) *ExcuseService {
	return &ExcuseService{db: db}
}

func (s *ExcuseService) GetRandomExcuse() (string, error) {
	query := `
		SELECT text 
		FROM excuses 
		WHERE active = true 
		ORDER BY RANDOM() 
		LIMIT 1
	`

	var text string
	err := s.db.QueryRow(context.Background(), query).Scan(&text)
	if err != nil {
		// Fallback to hardcoded excuses if database is empty
		return s.getFallbackExcuse(), nil
	}

	return text, nil
}

func (s *ExcuseService) CreateExcuse(req *models.CreateExcuseRequest) (*models.Excuse, error) {
	query := `
		INSERT INTO excuses (id, text, category, locale, source, quality_score, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, text, category, locale, source, quality_score, active, created_at, updated_at
	`

	now := time.Now()
	id := fmt.Sprintf("excuse_%d", time.Now().UnixNano()) // simple ID for now

	excuse := &models.Excuse{}
	err := s.db.QueryRow(context.Background(), query,
		id, req.Text, req.Category, req.Locale, req.Source, req.QualityScore, true, now, now,
	).Scan(
		&excuse.ID, &excuse.Text, &excuse.Category, &excuse.Locale,
		&excuse.Source, &excuse.QualityScore, &excuse.Active, &excuse.CreatedAt, &excuse.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create excuse: %w", err)
	}

	return excuse, nil
}

func (s *ExcuseService) getFallbackExcuse() string {
	fallbackExcuses := []string{
		"Sometimes you don't lose, you just take an L. Literally, database is not working lol",
		"The database is not working, so I'm just going to use a fallback excuse",
		"I'm just going to use a fallback excuse",
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return fallbackExcuses[rng.Intn(len(fallbackExcuses))]
}
