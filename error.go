package errors

import (
	"encoding/json"
)

type Error struct {
	Code    ErrorCode `json:"error_code"`
	Message string    `json:"error_message"`
	Debug   string    `json:"debug_messag,omitempty"`
}

var ErrUnknown = Error{}

func New(code ErrorCode, data ...string) Error {
	var errorMessage, debugMessage string
	switch len(data) {
	case 1:
		errorMessage = data[0]
	case 2:
		errorMessage = data[0]
		debugMessage = data[1]
	}
	err := Error{
		Code:    code,
		Message: errorMessage,
		Debug:   debugMessage,
	}
	return err
}

func ParseError(e error) Error {
	err := Error{}
	json.Unmarshal([]byte(e.Error()), &err)
	if err.Exception() {
		return ErrUnknown
	}
	return err
}

func (e Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e Error) ErrorCode() int64 {
	return int64(e.Code)
}

func (e Error) ErrorMessage() string {
	return e.Message
}

func (e Error) DebugMessage() string {
	return e.Debug
}

func (e Error) WithDebugMessage(message string) Error {
	e.Debug = message
	return e
}

func (e Error) Succ() bool {
	if e.Code == Succ {
		return true
	}
	return false
}

func (e Error) Exception() bool {
	return !e.Succ()
}
