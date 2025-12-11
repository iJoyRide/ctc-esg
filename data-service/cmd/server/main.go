package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iJoyRide/ctc-esg/data-service/internal/api/health"
	"github.com/iJoyRide/ctc-esg/data-service/internal/config"
	"github.com/iJoyRide/ctc-esg/data-service/internal/ingestion/mqtt"
)

func main() {

	cfg := config.Load()
	mqttService := mqtt.NewMQTTService(cfg)
	if err := mqttService.Init(mqttService.HandleSensorData); err != nil {
		log.Fatalf("[MQTT] Failed to initialize: %v", err)
	}

	r := gin.Default()

	health.RegisterRoutes(r.Group("/health"))

	log.Fatal(r.Run(":" + cfg.Port))
}
