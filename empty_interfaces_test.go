package gsl

import (
	"reflect"
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	type foo string
	type bar struct {
		a int
		b any
	}

	fooString := foo("foo")
	var v interface{}

	v = fooString
	testInterfaceTo[string](t, v, true)
	testInterfaceTo[[]byte](t, v, true)
	testInterfaceTo[bar](t, v, false)

	i := int8(69)
	v = i
	testInterfaceTo[float32](t, v, true)
	testInterfaceTo[string](t, v, true) // Numbers get converted into ASCII char byte

	v = int64(2000000)
	testInterfaceTo[string](t, v, true)
	testInterfaceTo[bar](t, v, false)

	v = bar{a: 69, b: false}
	testInterfaceTo[string](t, v, false)

	s := "henlo"
	testCompareInterfaceValues[string](t, s, fooString, true, false)

	s = string(fooString)
	testCompareInterfaceValues[foo](t, s, fooString, true, true)

	testCompareInterfaceValues[string](t, float64(512), fooString, false, false)
	testCompareInterfaceValues[string](t, uint8(69), fooString, true, false)
	testCompareInterfaceValues[string](t, int(512), fooString, true, false)

}

func testInterfaceTo[T any](t *testing.T, a interface{}, shouldPass bool) {
	v, err := InterfaceTo[T](a)
	log := func() {
		typeA := reflect.TypeOf(a).String()
		typeV := reflect.TypeOf(v).String()
		t.Logf("a: %v (%s), v: %v (%s)\n", a, typeA, v, typeV)
	}

	if shouldPass {
		if err != nil {
			log()
			t.Fatalf("expecting nil error, got %s\n", err.Error())
		}
	} else {
		if err == nil {
			log()
			t.Fatal("expecting non-nil error")
		}
	}
}

func testCompareInterfaceValues[T comparable](
	t *testing.T,
	a, b interface{},
	convertible, isEqual bool,
) {
	ok, err := CompareInterfaceValues[T](a, b)

	log := func() {
		t.Logf("a: %v, b: %v, convertible: %v, isEqual: %v\n", a, b, convertible, isEqual)
		t.Logf("ok: %v, err: %v\n", ok, err)
	}

	if convertible {
		if err != nil {
			log()
			t.Errorf("expecting nil error, got %s\n", err.Error())
		}
	} else {
		if err == nil {
			log()
			t.Error("expecting non-nil error, got nil")
		}
	}

	if isEqual != ok {
		log()
		t.Fatal("unexpected result")
	}
}
