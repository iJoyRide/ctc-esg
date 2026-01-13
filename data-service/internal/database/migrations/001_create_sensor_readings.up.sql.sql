CREATE TABLE IF NOT EXISTS sensor_readings (
  timestamp TIMESTAMPTZ NOT NULL,
  sensor_id TEXT NOT NULL,
  sensor TEXT NOT NULL,
  value DOUBLE PRECISION NOT NULL,
  PRIMARY KEY (timestamp, sensor_id)
);

SELECT create_hypertable(
  'sensor_readings',
  'timestamp',
  if_not_exists => TRUE
);
