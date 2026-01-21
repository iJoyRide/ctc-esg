package server

import "github.com/iJoyRide/ctc-esg/data-service/internal/server/handlers"

func (s *Server) RegisterRoutes() {
	s.router.GET("/gin_health", handlers.HealthCheck)

	readingsHandler := handlers.NewReadingsHandler(s.repo)
	api := s.router.Group("/api/v1")
	{
		api.GET("/readings/bucket", readingsHandler.GetBucketedReadings)
	}
}
