package main

import (
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"

	"roger/pkg/model"
	"roger/svc/router"
	"roger/svc/service"
)

const (
	name = "roger"
	port = 9933
)

func main() {
	var err error
	defer func() {
		fmt.Println(err)
	}()

	// uncomment below when release
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	var smsCredential = model.SMSCredential{
		SecretId:      "AKIDe6IEtHNCl22lEQLzNUKs99sJTN0LPDjy",
		SecretKey:     "0pNdcKnzuXuzoLdeOr3ycy3rOecJjskK",
		RequestMethod: "POST",
		Endpoint:      "sms.tencentcloudapi.com",
		SignMethod:    "HmacSHA1",
		Region:        "ap-guangzhou",
		SDKAppId:      "1400642103",
		SignName:      "eiicyn公众号",
	}
	svc := service.NewService(smsCredential)
	{
		// register service and routers
		group := server.Group(name)
		router.Register(svc, group)
	}

	if err = server.Run(fmt.Sprintf(":%d", port)); err != nil {
		return
	}
}
