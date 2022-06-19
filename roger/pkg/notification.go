package pkg

import (
	"encoding/json"
	"fmt"
	"roger/api"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"roger/pkg/model"
)

const (
	sendSMSByPhoneNumberErr = "Notification.SendSMSByPhoneNumberErr"
)

type NotificationService interface {
	SendSMSByPhoneNumber(phoneNumber model.PhoneNumber, template model.SMSTemplate) error
}

type notificationSvc struct {
	smsCredential model.SMSCredential
}

func NewNotificationSvc(smsCredential model.SMSCredential) NotificationService {
	svc := &notificationSvc{
		smsCredential: smsCredential,
	}
	return svc
}

func (n *notificationSvc) SendSMSByPhoneNumber(phoneNumber model.PhoneNumber, template model.SMSTemplate) error {
	credential := common.NewCredential(n.smsCredential.SecretId, n.smsCredential.SecretKey)

	// client profile object to set timeout and others
	cpf := profile.NewClientProfile()

	// `POST` by default
	cpf.HttpProfile.ReqMethod = n.smsCredential.RequestMethod
	cpf.HttpProfile.Endpoint = n.smsCredential.Endpoint
	cpf.SignMethod = n.smsCredential.SignMethod

	// set region
	client, _ := sms.NewClient(credential, n.smsCredential.Region, cpf)

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(n.smsCredential.SDKAppId)
	request.SignName = common.StringPtr(n.smsCredential.SignName)

	request.TemplateId = common.StringPtr("1323987")
	request.TemplateParamSet = common.StringPtrs(template.Content)
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber.String()})

	response, err := client.SendSms(request)

	// handle sdk exception
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("TencentCloudSDK error: %s\n", err)
		return api.NewError(sendSMSByPhoneNumberErr,
			fmt.Sprintf("fail to send sms by phone number: %s", phoneNumber.String()), err)
	}

	resp, _ := json.Marshal(response.Response)
	fmt.Printf("%s\n", resp)

	return nil
}
