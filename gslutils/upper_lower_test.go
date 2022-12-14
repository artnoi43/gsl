package gslutils

import "testing"

type stringAlias string

func testUpperLower[T ~string](
	t *testing.T,
	m map[T]T,
	f func(T) T,
) {
	for k, v := range m {
		if result := f(k); result != v {
			t.Logf("strings not matched:\nActual'%s' vs Expected'%s'", result, v)
			t.Fatal("strings not matched")
		}
	}
}

func TestToUpper(t *testing.T) {
	tests := map[stringAlias]stringAlias{
		"lol":  "LOL",
		"k_uy": "K_UY",
		"lA3":  "LA3",
	}
	testUpperLower(t, tests, ToUpper[stringAlias])
}

func TestToLower(t *testing.T) {
	tests := map[stringAlias]stringAlias{
		"lol":  "lol",
		"K_uY": "k_uy",
		"lA3%": "la3%",
	}
	testUpperLower(t, tests, ToLower[stringAlias])
}
