package gslutils

import "testing"

func TestInterfaceToT(t *testing.T) {
	i := 69
	var v interface{}
	v = i

	f, err := InterfaceTo[float64](v)
	if err != nil {
		t.Error(err.Error())
	}

	if f != float64(69) {
		t.Fatal("unexpected result")
	}
}
