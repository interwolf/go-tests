package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("5 numbers sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, array %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 5}, []int{2, 4})
	want := []int{9, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("sum the tails of normal slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3, 5}, []int{2, 4})
		want := []int{8, 4}
		checkSums(t, got, want)
	})

	t.Run("sum the tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)

	})

}
