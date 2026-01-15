package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iJoyRide/ctc-esg/data-service/internal/database"
	"github.com/iJoyRide/ctc-esg/data-service/internal/models"
)

type ReadingsHandler struct {
	repo *database.Repository
}

func NewReadingsHandler(repo *database.Repository) *ReadingsHandler {
	return &ReadingsHandler{repo: repo}
}

func (h *ReadingsHandler) GetBucketedReadings(c *gin.Context) {

	sensorID := c.Query("sensor_id")
	exist, err := h.repo.CheckSensorIdExists(c.Request.Context(), sensorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	} else if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "sensor not found"})
		return
	}

	
	startTime, err := time.Parse(time.RFC3339, c.Query("start_time"))
	if err != nil && c.Query("start_time") != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_time format. Use RFC3339 (e.g. 2006-01-02T15:04:05Z)"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, c.Query("end_time"))
	if err != nil && c.Query("end_time") != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_time format. Use RFC3339 (e.g. 2006-01-02T15:04:05Z)"})
		return 
	}

	req := models.GetReadingsRequest{
		BucketWidth: c.Query("bucket_width"),
		SensorID:    sensorID,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.repo.GetReadingsByBucket(c.Request.Context(), req.ToGetReadingsParams())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rows)
}