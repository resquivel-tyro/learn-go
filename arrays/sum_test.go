package arrays

import "testing"

func TestSumUsingArray(t *testing.T) {
	// can also initialise arrays like so:
	// numbersArray := [...]int{1, 2, 3, 4, 5}
	numbersArray := [5]int{1, 2, 3, 4, 5}

	got := SumArray(numbersArray)
	want := 15

	if got != want {
		// %v is the default format
		t.Errorf("got %d want %d given, %v", got, want, numbersArray)
	}
}

func TestSumUsingSlice(t *testing.T) {
	numbersSlice := []int{1, 2, 3}

	got := SumSlice(numbersSlice)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbersSlice)
	}
}