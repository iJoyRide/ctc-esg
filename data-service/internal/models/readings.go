package models

import (
	"errors"
	"time"

	db "github.com/iJoyRide/ctc-esg/data-service/internal/database/sqlc"
)

type GetReadingsRequest struct {
	BucketWidth string    `json:"bucket_width"`
    SensorID    string    `json:"sensor_id"`
    StartTime   time.Time `json:"start_time"`
    EndTime     time.Time `json:"end_time"`
}

func (g *GetReadingsRequest) Validate() error {
    if g.SensorID == "" {
        return errors.New("sensor_id field is required")
    }
    if !g.StartTime.IsZero() && !g.EndTime.IsZero() {
        if g.StartTime.After(g.EndTime) {
            return errors.New("start_time must be before end_time")
        }
    }
    return nil
}

func (g GetReadingsRequest) ToGetReadingsParams() db.GetReadingsByBucketParams {
	return db.GetReadingsByBucketParams{
		BucketWidth: g.BucketWidth,
		SensorID: g.SensorID,
		StartTime: g.StartTime,
		EndTime: g.EndTime,
	}
}