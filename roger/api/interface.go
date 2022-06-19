package api

import "roger/pkg/model"

type Service interface {
	HealthCheck() error
	SendSMSByPhoneNumber(phoneNumber model.PhoneNumber, template model.SMSTemplate) error
}
