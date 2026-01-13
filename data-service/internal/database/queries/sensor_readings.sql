-- name: InsertSensorReading :exec
INSERT INTO sensor_readings (timestamp, sensor, sensor_id, value)
VALUES ($1, $2, $3, $4);
