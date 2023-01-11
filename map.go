package gsl

// SliceFromMapValues collects map[K]V |m| into []V.
// If |m| is mil, SliceFromMapValues returns nil
func SliceFromMapValues[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}

	l := len(m)
	values := make([]V, l)

	var c int
	for _, v := range m {
		values[c] = v
		c++
	}

	return values
}

// SliceFromMapKeys collects map[K]V |m| into []K.
// If |m| is mil, SliceFromMapKeys returns nil.
func SliceFromMapKeys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}

	l := len(m)
	keys := make([]K, l)

	var c int
	for k := range m {
		keys[c] = k
		c++
	}

	return keys
}

// SlicesFromMap collects map[K]V |m| into ([]K, []V).
// If |m| is nil, SlicesFromMap returns nil.
func SlicesFromMap[K comparable, V any](m map[K]V) ([]K, []V) {
	if m == nil {
		return nil, nil
	}

	l := len(m)
	keys := make([]K, l)
	vals := make([]V, l)

	var c int
	for k, v := range m {
		keys[c] = k
		vals[c] = v
		c++
	}

	return keys, vals
}

// SliceFromMapValuesIf iterates over |m| and calls |f|
// with keys and values of |m|. For each true returned from |f|,
// the value in |m| gets appended to the return slice.
// If either |m| or |f| is nil, SliceFromMapKeysIf returns nil.
func SliceFromMapValuesIf[K comparable, V any](
	m map[K]V,
	f func(K, V) bool,
) []V {
	if m == nil || f == nil {
		return nil
	}

	var s []V
	for k, v := range m {
		if f(k, v) {
			s = append(s, v)
		}
	}

	return s
}

// SliceFromMapKeysIf iterates over |m| and calls |f|
// with keys and values of |m|. For each true returned from |f|,
// the key in |m| gets appended to the return slice.
// If either |m| or |f| is nil, SliceFromMapKeysIf returns nil.
func SliceFromMapKeysIf[K comparable, V any](
	m map[K]V,
	f func(K, V) bool,
) []K {
	if m == nil || f == nil {
		return nil
	}

	var s []K
	for k, v := range m {
		if f(k, v) {
			s = append(s, k)
		}
	}

	return s
}

// If either |m| or |f| is nil, SlicesFromMapIf returns nil.
func SlicesFromMapIf[K comparable, V any](m map[K]V, f func(K, V) bool) ([]K, []V) {
	if m == nil || f == nil {
		return nil, nil
	}

	var keys []K
	var vals []V
	for k, v := range m {
		if f(k, v) {
			keys = append(keys, k)
			vals = append(vals, v)
		}
	}

	return keys, vals
}
