package arraysnslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("5 numbers colection", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("\n Got: %d\n Want: %d\n Data: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSlicesDeep := func(t *testing.T, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\n Got: %v\n Want: %v", got, want)
		}
	}

	t.Run("Sum 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSlicesDeep(t, got, want)
	})

	t.Run("Sum 1 slice", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		want := []int{3}

		checkSlicesDeep(t, got, want)
	})
}
