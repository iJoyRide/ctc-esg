-- name: CheckSensorIdExists :one
SELECT EXISTS (
    SELECT 1 
    FROM sensor_readings 
    WHERE sensor_id = $1
);

-- name: GetReadingsByBucket :many
SELECT 
    time_bucket(sqlc.arg('bucket_width')::text::interval, timestamp)::timestamptz AS bucket,
    AVG(value)::float AS avg_value
FROM sensor_readings
WHERE sensor_id = sqlc.arg('sensor_id') 
  AND timestamp >= sqlc.arg('start_time') 
  AND timestamp <  sqlc.arg('end_time')
GROUP BY bucket
ORDER BY bucket ASC;

-- name: InsertSensorReading :exec
INSERT INTO sensor_readings (timestamp, sensor, sensor_id, value)
VALUES ($1, $2, $3, $4);