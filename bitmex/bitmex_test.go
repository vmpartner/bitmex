package bitmex

import (
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	message := []byte(`{"Success":true,"Subscribe":"","Request":"","Table":"order","Action":"partial","Data":""}`)
	res, err := DecodeMessage(message)
	if err != nil || res.Success != true || res.Table != "order" {
		t.Error("Error decode message")
	}
}
