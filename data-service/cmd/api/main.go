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

	dbService := database.NewDatabaseService(cfg)
	if err := dbService.Init(); err != nil {
		log.Fatalf("[Database] Failed to initialize: %v", err)
	}

	repo := database.NewRepository(dbService)

	mqttService := mqtt.NewMQTTService(cfg, repo)
	if err := mqttService.Init(mqttService.HandleSensorData); err != nil {
		log.Fatalf("[MQTT] Failed to initialize: %v", err)
	}

	srv := server.NewServer(mqttService, dbService, repo)
	log.Fatal(srv.Run(":" + cfg.Port))
}
