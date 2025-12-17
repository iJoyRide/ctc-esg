package database

import (
	"context"

	"github.com/iJoyRide/ctc-esg/data-service/internal/models"
)

func (d *DatabaseService) Insert(
	ctx context.Context,
	reading models.SensorReading,
) error {

	query := `
		INSERT INTO sensor_readings (timestamp, sensor, sensor_id, value)
		VALUES ($1, $2, $3, $4)
	`

	_, err := d.db.ExecContext(
		ctx,
		query,
		reading.Timestamp,
		reading.Sensor,
		reading.SensorID,
		reading.Value,
	)

	return err
}
