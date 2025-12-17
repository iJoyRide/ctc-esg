package database

import (
	"context"
	"log"
)

func (d *DatabaseService) createSchema(ctx context.Context) error {
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS sensor_readings (
            timestamp TIMESTAMPTZ NOT NULL,
            sensor_id TEXT NOT NULL,
            sensor TEXT,
            value DOUBLE PRECISION,
            PRIMARY KEY (timestamp, sensor_id)
        );
    `

	if _, err := d.db.ExecContext(ctx, createTableQuery); err != nil {
		return err
	}

	hyperTableQuery := `
        SELECT create_hypertable(
            'sensor_readings', 
            'timestamp',
            if_not_exists => TRUE
        );
    `

	if _, err := d.db.ExecContext(ctx, hyperTableQuery); err != nil {
		log.Printf("[Database] create_hypertable skipped or failed: %v", err)
	}

	log.Println("[Database] Schema verified/created")
	return nil
}
