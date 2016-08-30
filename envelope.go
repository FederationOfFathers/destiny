package destiny

import (
	"encoding/json"
	"fmt"
)

type dataEnvelope struct {
	Data *json.RawMessage `json:"data"`
}

func (d *dataEnvelope) into(into interface{}) error {
	if d == nil || d.Data == nil {
		return fmt.Errorf("envelope did not contain data")
	}
	return json.Unmarshal(*d.Data, &into)
}

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

func (e *envelope) data() (*dataEnvelope, error) {
	var rval = &dataEnvelope{}
	err := e.into(&rval)
	return rval, err
}

func (e *envelope) into(into interface{}) error {
	return json.Unmarshal(e.Response, &into)
}
