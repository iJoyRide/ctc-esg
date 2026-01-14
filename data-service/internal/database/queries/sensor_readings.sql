-- name: InsertSensorReading :exec
INSERT INTO sensor_readings (timestamp, sensor, sensor_id, value)
VALUES ($1, $2, $3, $4);

SELECT time_bucket('5 min', time) AS bucket 
FROM sensor_readings
WHERE timestamp BETWEEN $1 AND $2
ORDER BY timestamp ASC