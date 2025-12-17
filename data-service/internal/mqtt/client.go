package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iJoyRide/ctc-esg/data-service/internal/config"
)

type MQTTService struct {
	client mqtt.Client
	cfg    *config.Config
}

func NewMQTTService(configuration *config.Config) *MQTTService {
	return &MQTTService{cfg: configuration}
}

func (m *MQTTService) Init(handler mqtt.MessageHandler) error {
	opts := mqtt.NewClientOptions().
		AddBroker(m.cfg.MQTT.Broker).
		SetClientID(m.cfg.MQTT.ClientID).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(3 * time.Second)

	opts.OnConnect = func(c mqtt.Client) {
		c.Subscribe(m.cfg.MQTT.Topic, byte(m.cfg.MQTT.QOS), handler)
	}

	m.client = mqtt.NewClient(opts)
	token := m.client.Connect()
	token.Wait()
	log.Printf("[MQTT] initialised")
	return token.Error()

}
