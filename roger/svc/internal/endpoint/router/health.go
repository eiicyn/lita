package router

import (
	"github.com/gin-gonic/gin"
	"roger/api"
	"roger/svc/internal/endpoint/response"
)

func healthCheck(svc api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := svc.HealthCheck()
		c.JSON(200, response.HealthCheckResponse{
			RogerResponse: response.RogerResponse{Roger: err},
			Data:          "pong",
		})
	}
}

func RegisterHealthRouter(svc api.Service, g *gin.RouterGroup) {
	g.GET("/health", healthCheck(svc))
}
