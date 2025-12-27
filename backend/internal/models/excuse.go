package models

import (
	"time"
)

// Excuse represents a badminton excuse in the database
type Excuse struct {
	ID           string    `json:"id" db:"id"`
	Text         string    `json:"text" db:"text"`
	Category     string    `json:"category" db:"category"`
	Locale       string    `json:"locale" db:"locale"`
	Source       string    `json:"source" db:"source"`
	QualityScore int       `json:"quality_score" db:"quality_score"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// CreateExcuseRequest represents the request to create a new excuse
type CreateExcuseRequest struct {
	Text         string `json:"text" validate:"required"`
	Category     string `json:"category" validate:"required,oneof=singles doubles general mixed"`
	Locale       string `json:"locale"`
	Source       string `json:"source"`
	QualityScore int    `json:"quality_score"`
}

// ExcuseResponse represents the response for excuse operations
type ExcuseResponse struct {
	Message string  `json:"message"`
	Excuse  *Excuse `json:"excuse,omitempty"`
}
