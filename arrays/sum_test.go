package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("empty collection", func(t *testing.T) {
		var numbers []int

		got := Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of 5 integers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})


}
