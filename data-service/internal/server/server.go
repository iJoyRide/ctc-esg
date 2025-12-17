package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iJoyRide/ctc-esg/data-service/internal/database"
	"github.com/iJoyRide/ctc-esg/data-service/internal/mqtt"
)

type Server struct {
	router *gin.Engine
	mqtt   *mqtt.MQTTService
	db     *database.DatabaseService
}

func NewServer(mqttService *mqtt.MQTTService, dbService *database.DatabaseService) *Server {
	router := gin.Default()

	s := &Server{
		router: router,
		mqtt:   mqttService,
		db:     dbService,
	}

	s.RegisterRoutes()

	return s
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
