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
	rect := Rectangle{12.0, 6.0}
	got := Area(rect)
	want := 72.0

	if got != want {
		t.Errorf("\n Got: %.2f\n Want: %.2f", got, want)
	}
}
