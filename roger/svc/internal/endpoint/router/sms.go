package router

import (
	"github.com/gin-gonic/gin"

	"roger/api"
	"roger/svc/internal/endpoint/request"
	"roger/svc/internal/endpoint/response"
)

func sendSMSByPhoneNumber(svc api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := request.SendSMSByPhoneNumberRequest{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(200, response.SendSMSByPhoneNumberResponse{
				RogerResponse: response.RogerResponse{Roger: api.JSONUnmarshalErr(err)},
			})
			return
		}
		err := svc.SendSMSByPhoneNumber(req.PhoneNumber, req.Template)
		c.JSON(200, response.SendSMSByPhoneNumberResponse{
			RogerResponse: response.RogerResponse{Roger: err},
		})
	}
}

func RegisterSMSRouter(svc api.Service, g *gin.RouterGroup) {
	g.POST("/sms", sendSMSByPhoneNumber(svc))
}
