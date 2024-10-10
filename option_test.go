package gsl_test

import (
	"testing"

	"github.com/soyart/gsl"
)

func TestOption(t *testing.T) {
	s := gsl.OptionSome("some")

	if gsl.OptionIsNone(s) {
		t.Fatalf("unexpected none option")
	}

	if !gsl.OptionIsSome(s) {
		t.Fatalf("unexpected not-some option")
	}

	s = gsl.OptionNone[string]()

	if !gsl.OptionIsNone(s) {
		t.Fatalf("unexpected not-none option")
	}

	if gsl.OptionIsSome(s) {
		t.Fatalf("unexpected some option")
	}
}
