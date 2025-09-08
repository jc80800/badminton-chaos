package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseService struct {
	pool *pgxpool.Pool
}

func NewDatabaseService() (*DatabaseService, error) {
	connStr := "postgres://badminton_user:badminton_pass@localhost:5432/badminton_chaos?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connected successfully")

	return &DatabaseService{pool: pool}, nil
}

func (db *DatabaseService) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}

func (db *DatabaseService) GetPool() *pgxpool.Pool {
	return db.pool
}
