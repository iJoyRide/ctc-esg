#!/bin/sh

CONFIG_FILE=/mosquitto/config/mosquitto.conf

# Check if config exists, if not create a default one
if [ ! -f "$CONFIG_FILE" ]; then
    echo "mosquitto.conf not found, creating default config..."
    cat <<EOL > "$CONFIG_FILE"
listener 1883
allow_anonymous true
persistence true
persistence_location /mosquitto/data/
log_dest file /mosquitto/log/mosquitto.log
EOL
fi

echo "mosquitto.conf already exists"
echo "Mosquitto broker started with config file"

# Start Mosquitto
exec mosquitto -c "$CONFIG_FILE"

