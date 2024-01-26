package gsl

import "testing"

func TestIsWellFormed(t *testing.T) {
	shouldOk := []string{
		"",
		"[]",
		"()",
		"foo[(foo)]",
		"([bar(foo)])",
		"[()[{}]]",
		"[()[{}]]<>",
		"[()[{}]]<[([])]>",
		"foo(bar[baz])",
	}

	shouldErr := []string{
		"[",
		"(",
		"[)",
		"[[(]]",
		"foo[(]",
	}

	for i, s := range shouldOk {
		err := IsWellClosed(s)
		if err != nil {
			t.Fatalf("unexpected error for string #%d '%s': %s", i, s, err.Error())
		}
	}

	for i, s := range shouldErr {
		err := IsWellClosed(s)
		if err == nil {
			t.Fatalf("unexpected ok result for bad string #%d '%s'", i, s)
		}
	}
}
