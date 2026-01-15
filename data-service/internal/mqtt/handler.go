package mqtt

import (
	"context"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iJoyRide/ctc-esg/data-service/internal/models"
)

// HandleSensorData is called whenever a message arrives on the subscribed topic
func (m *MQTTService) HandleSensorData(_ mqtt.Client, msg mqtt.Message) {
	var payload models.SensorReadingPayload

	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Printf("[MQTT] Invalid payload: %v", err)
		return
	}

	if err := payload.Validate(); err != nil {
		log.Printf("[MQTT] Invalid sensor reading: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.repo.InsertSensorReading(ctx, payload.ToInsertParams()); err != nil {
		log.Printf("[DB] insert failed: %v", err)
		return
	}

	log.Printf("[MQTT] Received sensor reading: %+v", payload)
}
