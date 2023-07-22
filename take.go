package ranger

// Take wraps a push iterator with a function that will either take the
// given integer number of values and stop, or stop when the underlying
// sequence stops.
//
// No special treatment for zero or negative numbers, both of which will
// empty the sequence.
func Take[K, V any](num int, f func(func(K, V) bool) bool) func(func(K, V) bool) bool {
	return func(yield func(key K, val V) bool) bool {
		if num <= 0 {
			return false
		}
		for k, v := range f {
			if !yield(k, v) {
				return false
			}
			num--
			if num <= 0 {
				return false
			}
		}
		return false
	}
}

// IntRange turns a range on an int into a component that can be used in
// functional range pipelines.
func IntRange(r int) func(func(int, int)bool)bool {
	return func(yield func(int, int)bool) bool {
		for i := range r {
			if !yield(i, i) {
				return false
			}
		}
		return false
        }
}
