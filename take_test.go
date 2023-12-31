package ranger

import "testing"
import "reflect"

func TestTake(t *testing.T) {
	s := []int{}
	for i  := range Take(5, IntRange(25)) {
		s = append(s, i)
	}
	if !reflect.DeepEqual(s, []int{0, 1, 2, 3, 4}) {
		t.Fatal("Take does not work")
	}
}

type rangeTest struct {
	Start int
	Stop int
	Jump int
	Result []int
}

func TestRangeOver(t *testing.T) {
	for idx, test := range []rangeTest{
		{0, 2, 1, []int{0, 1}},
		{0, -1, 8, []int{}},
		{0, 1, 1, []int{0}},
		{0, 10, 2, []int{0, 2, 4, 6, 8}},
	}{
		res := []int{}
		for i := range RangeOver(test.Start, test.Stop, test.Jump) {
			res = append(res, i)
		}
		if !reflect.DeepEqual(res, test.Result) {
			t.Fatalf("in test %d, expected: %v, got: %v",
				idx, test.Result, res)
		}
	}
}

// This tries to benchmark a basic Take 5 from a 25 int range. Due to the
// simplicity of all the operations this should give a good view into the
// overhead of the wrapper versus doing it directly.
//
// Don't overinterpret this benchmark; this is a test of the overhead. Real
// loops with real payload would ideally do more.
func BenchmarkTakeAsForLoop(b *testing.B) {
	s := make([]int, 5)
	for range b.N {
		take := 5
		INNER: for i := range 25 {
			// need to do "something" with the loop to compare
			// to Take fairly
			s[i] = i
			take--
			if take <= 0 {
				break INNER
			}
		}
	}
}

func BenchmarkTakeViaRangeComposition(b *testing.B) {
	s := make([]int, 5)
	for range b.N {
		for i := range Take(5, IntRange(25)) {
			s[i] = i
		}
	}
}

func BenchmarkRangeDirectly(b *testing.B) {
	for range b.N {
		for range 25 {
		}
	}
}

func BenchmarkRangeFunc(b *testing.B) {
	for range b.N {
		for range IntRange(25) {
		}
	}
}
