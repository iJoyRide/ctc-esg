package database

import (
	"context"

	db "github.com/iJoyRide/ctc-esg/data-service/internal/database/sqlc"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(dbService *DatabaseService) *Repository {
	return &Repository{
		q: dbService.Queries(),
	}
}

func (r *Repository) Insert(ctx context.Context, params db.InsertSensorReadingParams) error {
	return r.q.InsertSensorReading(ctx, params)
}
