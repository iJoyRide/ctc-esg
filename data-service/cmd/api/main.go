package main

import (
	"log"

	"github.com/iJoyRide/ctc-esg/data-service/internal/config"
	"github.com/iJoyRide/ctc-esg/data-service/internal/database"
	"github.com/iJoyRide/ctc-esg/data-service/internal/mqtt"
	"github.com/iJoyRide/ctc-esg/data-service/internal/server"
)

func main() {
	cfg := config.Load()

	// 1. Initialize DATABASE FIRST
	dbService := database.NewDatabaseService(cfg)
	if err := dbService.Init(); err != nil {
		log.Fatalf("[Database] Failed to initialize: %v", err)
	}

	// 2. Then initialize MQTT with database
	mqttService := mqtt.NewMQTTService(cfg, dbService) // ← Pass dbService here
	if err := mqttService.Init(mqttService.HandleSensorData); err != nil {
		log.Fatalf("[MQTT] Failed to initialize: %v", err)
	}

	// 3. Start server
	srv := server.NewServer(mqttService, dbService)
	log.Fatal(srv.Run(":" + cfg.Port))
}
