package server

import "github.com/iJoyRide/ctc-esg/data-service/internal/server/handlers"

func (s *Server) RegisterRoutes() {
	s.router.GET("/health", handlers.HealthCheck)
}
