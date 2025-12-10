package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iJoyRide/ctc-esg/data-service/config"
	"github.com/iJoyRide/ctc-esg/data-service/internal/api/health"
)

func main() {

	cfg := config.Load()

	r := gin.Default()

	health.RegisterRoutes(r.Group("/health"))

	log.Fatal(r.Run(":" + cfg.Port))
}
