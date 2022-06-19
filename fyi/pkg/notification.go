package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fyi/pkg/model"
	"net/http"
)

type sendSMSRequest struct {
	PhoneNumber model.PhoneNumber `json:"phoneNumber"`
	Template    model.SMSTemplate `json:"template"`
}

func SendSMS(content string) {
	req := sendSMSRequest{
		PhoneNumber: model.PhoneNumber{
			CountryCode: rogerPhoneNumberCountryCode,
			Number:      rogerPhoneNumber,
			Extension:   "",
		},
		Template: model.SMSTemplate{
			Name:    name,
			Content: []string{fmt.Sprintf("领导想%s", content), "10", ""},
		},
	}
	body, _ := json.Marshal(req)

	_, err := http.Post(
		fmt.Sprintf("%s/roger/alorel/andethoras/ethil/sms", rogerServerAddress),
		model.HTTPPostRequestContentType,
		bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Error occurs: %e\nConnect in wechat please X(\n", err)
		return
	}

	fmt.Println("Roger that!")
}
