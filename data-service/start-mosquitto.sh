#!/bin/sh

CONFIG_FILE=/mosquitto/config/mosquitto.conf

echo "Ensuring mosquitto.conf contains listener config..."

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Creating mosquitto.conf..."
    cat <<EOL > "$CONFIG_FILE"
listener 1883 0.0.0.0
allow_anonymous true
persistence true
persistence_location /mosquitto/data/
log_dest stdout
log_type warning
EOL
else
    if ! grep -q "listener 1883" "$CONFIG_FILE"; then
        echo "Adding missing listener..."
        echo "listener 1883 0.0.0.0" >> "$CONFIG_FILE"
        echo "allow_anonymous true" >> "$CONFIG_FILE"
    fi

    if ! grep -q "log_type" "$CONFIG_FILE"; then
        echo "log_type warning" >> "$CONFIG_FILE"
    fi

    echo "Using existing mosquitto.conf"
fi

echo "Mosquitto starting..."

exec mosquitto -c "$CONFIG_FILE"
