package smi

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("\n Got: %.2f\n Want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("\n Got: %g\n Want: %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		checkArea(t, rect, 72.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
