package gomaith

// S is a convenience function that returns a slice of its arguments.
func S[T any](ts ...T) []T {
	return ts
}

// M is a convenience function that returns a map-set of the its arguments.
func M[T comparable](ts ...T) map[T]struct{} {
	m := make(map[T]struct{})
	for _, t := range ts {
		m[t] = struct{}{}
	}
	return m
}

// difference returns a map consisting of the keys in m1 that do not exist in m2
// and vice-versa.
func difference[T comparable](m1, m2 map[T]struct{}) map[T]struct{} {
	d := make(map[T]struct{})
	for k, _ := range m1 {
		if _, found := m2[k]; !found {
			d[k] = struct{}{}
		}
	}
	for k, _ := range m2 {
		if _, found := m1[k]; !found {
			d[k] = struct{}{}
		}
	}
	return d
}

// lastIndex returns last index of the item in t satisfying fn.
func lastIndex[T comparable](ts []T, fn func(T) bool) int {
	for i := len(ts) - 1; i >= 0; i-- {
		if fn(ts[i]) {
			return i
		}
	}
	return -1
}
