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
	repo *database.Repository
}

func NewServer(mqttService *mqtt.MQTTService, dbService *database.DatabaseService, repoService *database.Repository) *Server {
	router := gin.Default()

	s := &Server{
		router: router,
		mqtt:   mqttService,
		db:     dbService,
		repo: repoService,
	}

	s.RegisterRoutes()

	return s
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
