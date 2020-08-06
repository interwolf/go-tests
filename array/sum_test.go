package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("5 numbers sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, array %v", got, want, numbers)
		}
	})

	// t.Run("any numbers sum", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d, want %d, array %v", got, want, numbers)
	// 	}
	// })

}
