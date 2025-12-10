package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string

	MQTT struct {
		Broker   string
		ClientID string
		Topic    string
		QOS      int
	}
}

func Load() *Config {
	viper.AutomaticEnv()

	cfg := &Config{}

	cfg.Port = viper.GetString("PORT")
	cfg.MQTT.Broker = viper.GetString("MQTT_BROKER")
	cfg.MQTT.ClientID = viper.GetString("MQTT_CLIENT_ID")
	cfg.MQTT.Topic = viper.GetString("MQTT_TOPIC")
	cfg.MQTT.QOS = viper.GetInt("MQTT_QOS")

	validate(cfg)

	return cfg
}

func validate(cfg *Config) {
	required := map[string]string{
		"PORT":        cfg.Port,
		"MQTT_BROKER": cfg.MQTT.Broker,
		"MQTT_TOPIC":  cfg.MQTT.Topic,
	}

	for key, val := range required {
		if val == "" {
			log.Fatalf("FATAL: Required environment variable %s is missing!", key)
		}
	}
}
