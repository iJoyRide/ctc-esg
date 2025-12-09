package health

import "github.com/gin-gonic/gin"

func RegisterRoutes(router_group *gin.RouterGroup) {
	router_group.GET("/", HealthCheck)
}
