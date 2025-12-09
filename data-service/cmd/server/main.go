package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iJoyRide/ctc-esg/data-service/internal/api/health"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("FATAL: Required environment variable PORT is not set. Service cannot start.")
	}

	router := gin.Default()

	health.RegisterRoutes(router.Group("/health"))

	log.Fatal(router.Run(":" + port))
}
