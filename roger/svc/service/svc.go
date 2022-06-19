package service

import (
	"roger/api"
	"roger/pkg"
	"roger/pkg/model"
)

type service struct {
	notificationSvc pkg.NotificationService
}

func NewService(smsCredential model.SMSCredential) api.Service {
	svc := &service{
		notificationSvc: pkg.NewNotificationSvc(smsCredential),
	}
	return svc
}

func (s *service) HealthCheck() error {
	return api.Success()
}

func (s *service) SendSMSByPhoneNumber(phoneNumber model.PhoneNumber, template model.SMSTemplate) error {
	if err := s.notificationSvc.SendSMSByPhoneNumber(phoneNumber, template); err != nil {
		return err
	}

	return api.Success()
}
