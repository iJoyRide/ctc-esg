package models

import "time"

type SensorReading struct {
	Timestamp time.Time `json:"timestamp"`
	Sensor    string    `json:"sensor"`
	Value     float64   `json:"value"`
}
