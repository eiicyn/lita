package request

import (
	"roger/pkg/model"
)

type SendSMSByPhoneNumberRequest struct {
	PhoneNumber model.PhoneNumber `json:"phoneNumber"`
	Template    model.SMSTemplate `json:"template"`
}
