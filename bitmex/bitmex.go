package bitmex

import (
	"encoding/json"
)

type Response struct {
	Success   bool            `json:"success,omitempty"`
	Subscribe string          `json:"subscribe,omitempty"`
	Request   interface{}     `json:"request,omitempty"`
	Table     string          `json:"table,omitempty"`
	Action    string          `json:"action,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

func DecodeMessage(message []byte) (Response, error) {
	var res Response
	err := json.Unmarshal(message, &res)

	return res, err
}
