package mqtt

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iJoyRide/ctc-esg/data-service/internal/models"
)

// HandleSensorData is called whenever a message arrives on the subscribed topic
func (m *MQTTService) HandleSensorData(_ mqtt.Client, msg mqtt.Message) {
	var reading models.SensorReading

	if err := json.Unmarshal(msg.Payload(), &reading); err != nil {
		log.Printf("[MQTT] Invalid payload: %v", err)
		return
	}

	log.Printf("[MQTT] Received sensor reading: %+v", reading)
}
