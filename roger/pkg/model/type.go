package model

import (
	"fmt"
)

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
	Extension   string `json:"extension"`
}

func (p PhoneNumber) String() string {
	return fmt.Sprintf("+%s%s%s", p.CountryCode, p.Number, p.Extension)
}

type SMSTemplate struct {
	Id      string
	Name    string   `json:"name"`
	Content []string `json:"content"`
}

// SMSCredential used tencent cloud service currently, so it's customized
type SMSCredential struct {
	SecretId      string
	SecretKey     string
	RequestMethod string
	Endpoint      string
	SignMethod    string
	Region        string
	SDKAppId      string
	SignName      string
}
