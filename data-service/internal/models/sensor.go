package models

import (
	"errors"
	"time"
)

type SensorReading struct {
	Timestamp time.Time `json:"timestamp"`
	Sensor    string    `json:"sensor"`
	SensorID  string    `json:"sensor_id"`
	Value     float64   `json:"value"`
}

func (s *SensorReading) Validate() error {
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
