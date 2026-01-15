package database

import (
	"context"

	db "github.com/iJoyRide/ctc-esg/data-service/internal/database/sqlc"
)

type SensorRepository interface {
    InsertSensorReading(ctx context.Context, params db.InsertSensorReadingParams) error
    GetReadingsByBucket(ctx context.Context, params db.GetReadingsByBucketParams) ([]db.GetReadingsByBucketRow, error)
    CheckSensorIdExists(ctx context.Context, sensorID string) (bool, error)
}

type Repository struct {
	q *db.Queries
}

func NewRepository(dbService *DatabaseService) *Repository {
	return &Repository{
		q: dbService.Queries(),
	}
}

func (r *Repository) InsertSensorReading(ctx context.Context, params db.InsertSensorReadingParams) error {
	return r.q.InsertSensorReading(ctx, params)
}

func (r *Repository) GetReadingsByBucket(ctx context.Context, params db.GetReadingsByBucketParams) ([]db.GetReadingsByBucketRow, error) {
	return r.q.GetReadingsByBucket(ctx, params)
}

func (r *Repository) CheckSensorIdExists(ctx context.Context, sensorID string) (bool, error) {
    return r.q.CheckSensorIdExists(ctx, sensorID)
}