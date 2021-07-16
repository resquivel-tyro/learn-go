package arrays

import "testing"

func TestSum(t *testing.T) {
	// can also initialise arrays like so:
	// numbers := [...]int{1, 2, 3, 4, 5}
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		// %v is the default format
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}