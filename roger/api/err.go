package api

import "fmt"

type RogerError struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Err     error       `json:"error"`
}

const (
	RogerSuccess          = "ROGER-00000000"
	RogerJSONMarshalErr   = "ROGER-00000001"
	RogerJSONUnmarshalErr = "ROGER-00000002"
	RogerBadRequestErr    = "ROGER-00000400"
	RogerUnauthorizedErr  = "ROGER-00000401"
	RogerForbiddenErr     = "ROGER-00000403"
	RogerInternalErr      = "ROGER-00000500"
)

func NewError(code string, message string, err error) error {
	return &RogerError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *RogerError) Error() string {
	return fmt.Sprintf("[RogerError] Code=%s, Message=%s, Error=%s", e.Code, e.Message, e.Err.Error())
}

func Success() error {
	return &RogerError{
		Code:    RogerSuccess,
		Message: "Success.",
		Err:     nil,
	}
}

func JSONUnmarshalErr(err error) error {
	return &RogerError{
		Code:    RogerJSONUnmarshalErr,
		Message: "json unmarshal error",
		Err:     err,
	}
}
