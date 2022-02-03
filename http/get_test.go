package http

import (
	"testing"
)

func TestGetAndParse(t *testing.T) {
	var c struct {
		Count int `json:"count"`
	}
	if err := GetAndParse("https://limit-orders.1inch.io/v2.0/1/limit-order/count", &c); err != nil {
		t.Error("Error from GetAndParse", err.Error())
	}
}
