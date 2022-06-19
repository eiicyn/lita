package model

import (
	"fmt"
)

const HTTPPostRequestContentType = "application/json"

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

type CmdDict struct {
	Cmd     string
	Chinese string
	English string
}
