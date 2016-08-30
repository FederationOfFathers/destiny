package destiny

import (
	"encoding/json"
	"fmt"
)

type envelope struct {
	ErrorCode       int
	ErrorStatus     string
	Message         string
	MessageData     json.RawMessage
	Response        json.RawMessage
	ThrottleSeconds json.Number
}

func (e *envelope) Error() string {
	return fmt.Sprintf("%d:%s -- %s", e.ErrorCode, e.ErrorStatus, e.Message)
}

func (e *envelope) success() bool {
	return e.ErrorCode == 1
}
