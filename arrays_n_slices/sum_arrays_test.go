package arraysnslices

import "testing"

func TestSum(t *testing.T) {
	assertion := func(t *testing.T, got int, want int, data []int) {
		if got != want {
			t.Errorf("\n Got: %d\n Want: %d\n Data: %v", got, want, data)
		}
	}

	t.Run("5 numbers colection", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertion(t, got, want, numbers)
	})
}
