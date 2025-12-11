package models

import "time"

type SensorReading struct {
	Timestamp time.Time `json:"timestamp"`
	Sensor    string    `json:"sensor"`
	SensorID  string    `json:"sensor_id"`
	Value     float64   `json:"value"`
}
