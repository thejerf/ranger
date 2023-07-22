package ranger

import "testing"
import "fmt"
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
