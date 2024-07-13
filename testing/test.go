package testing

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of numbers in array", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}

		got := Sum(numbers)
		want := 16

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})
}

func Sum(list []int) int {
	var sum int = 0
	for _, i := range list {
		sum += i
	}
	return sum
}
