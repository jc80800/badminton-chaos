# Badminton Chaos - Development Commands

.PHONY: help db-up db-down db-reset backend frontend clean

help:
	@echo "db-up, db-down, db-reset, backend, frontend, clean"

db-up:
	@echo "Starting database..."
	docker compose up -d postgres

db-down:
	@echo "Stopping database..."
	docker compose down

db-reset:
	@echo "Resetting database..."
	docker compose down -v
	docker compose up -d postgres

backend:
	@echo "Starting backend..."
	cd backend && go run cmd/main.go

frontend:
	@echo "Starting frontend..."
	cd frontend && npm run dev

clean:
	@echo "Cleaning up..."
	docker compose down -v
	docker volume prune -f
