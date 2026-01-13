package models

import (
	"errors"
	"time"

	db "github.com/iJoyRide/ctc-esg/data-service/internal/database/sqlc"
)

type SensorReadingPayload struct {
	Timestamp time.Time `json:"timestamp"`
	Sensor    string    `json:"sensor"`
	SensorID  string    `json:"sensor_id"`
	Value     float64   `json:"value"`
}

func (s *SensorReadingPayload) Validate() error {
	if s.Sensor == "" {
		return errors.New("sensor field is required")
	}
	if s.SensorID == "" {
		return errors.New("sensor_id field is required")
	}
	if s.Timestamp.IsZero() {
		return errors.New("timestamp is required")
	}
	return nil
}

func (s SensorReadingPayload) ToInsertParams() db.InsertSensorReadingParams {
	return db.InsertSensorReadingParams{
		Timestamp: s.Timestamp,
		Sensor:    s.Sensor,
		SensorID:  s.SensorID,
		Value:     s.Value,
	}
}
