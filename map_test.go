package gsl

import (
	"reflect"
	"testing"
)

func TestSliceFromMap(t *testing.T) {
	m0 := map[uint8]string{}
	s0 := SliceFromMapValues(m0)

	if s0 == nil {
		t.Fatal("expecting 0-length slice, not nil")
	}

	s1 := SliceFromMapValues[string, string](nil)
	if s1 != nil {
		t.Fatal("expecting nil slice")
	}

	m2 := map[uint8]string{
		1: "one",
		2: "two",
		4: "four",
	}

	keys := SliceFromMapKeys(m2)
	vals := SliceFromMapValues(m2)

	var seenKeys = make(map[uint8]bool)
	var seenVals = make(map[string]bool)

	reset := func() {
		for k, v := range m2 {
			seenKeys[k], seenVals[v] = false, false
		}
	}

	check := func() {
		for _, key := range keys {
			seenKeys[key] = true
		}

		for _, val := range vals {
			seenVals[val] = true
		}

		for key, seen := range seenKeys {
			if !seen {
				t.Fatal("missing a key", key)
			}
		}

		for val, seen := range seenVals {
			if !seen {
				t.Fatal("missing a value", val)
			}
		}
	}

	reset()
	check()

	reset()
	keys, vals = SlicesFromMap(m2)
	check()
}

func TestSliceFromMapIf(t *testing.T) {
	m := map[uint64]string{
		1:   "one",
		2:   "two",
		3:   "three",
		69:  "foo",
		100: "cent",
	}

	filterEven := func(i uint64, _ string) bool {
		return i%2 == 0
	}

	// Filter for even number key
	resultKeys := SliceFromMapKeysIf(m, filterEven)
	if l := len(resultKeys); l != 2 {
		t.Fatalf("unexpected resultKeys length, expecting 2, got %d", l)
	}
	if !reflect.DeepEqual(resultKeys, []uint64{2, 100}) {
		if !reflect.DeepEqual(resultKeys, []uint64{100, 2}) {
			t.Fatalf("unexpected resultKeys values, expecting [2, 100] or [100, 2], got %v", resultKeys)
		}
	}

	resultVals := SliceFromMapValuesIf(m, filterEven)
	if l := len(resultKeys); l != 2 {
		t.Fatalf("unexpected resultKeys length, expecting 2, got %d", l)
	}
	if !reflect.DeepEqual(resultVals, []string{"two", "cent"}) {
		if !reflect.DeepEqual(resultVals, []string{"cent", "two"}) {
			t.Fatalf("unexpected resultKeys values, expecting [\"two\", \"cent\"] [\"cent\", \"two\"], got %v", resultKeys)
		}
	}

	checkLens := func() {
		if lKeys, lVals := len(resultKeys), len(resultVals); lKeys != lVals {
			t.Fatalf("resultKeys length %d not matched resultVals length %d", lKeys, lVals)
		}
	}

	checkLens()
	resultKeys, resultVals = SlicesFromMapIf(m, filterEven)
	checkLens()

	for i, key := range resultKeys {
		v, ok := m[key]
		if !ok {
			t.Fatalf("key %v not found in map %+v", key, m)
		}
		val := resultVals[i]

		if v != val {
			t.Log(resultKeys, resultVals)
			t.Fatalf("bad value collected, expecting %v, got %v", v, val)
		}
	}
}
