package router

import (
	"github.com/gin-gonic/gin"
	"roger/api"

	"roger/svc/internal/endpoint/router"
)

const (
	secretUrl = "alorel/andethoras/ethil"
)

func Register(svc api.Service, group *gin.RouterGroup) {
	v1 := group.Group(secretUrl)
	router.RegisterHealthRouter(svc, v1)
	router.RegisterSMSRouter(svc, v1)
}
