package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertCorrect := func(t *testing.T, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func (t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		assertCorrect(t, got, want)
	})

	t.Run("safely sum empty splices", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{0,9})
		want := []int{0,9}

		assertCorrect(t, got, want)
	})
}