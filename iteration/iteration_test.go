package iteration

import "testing"

func TestRepetir(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("\n Want: '%s' \n Got: '%s'", got, want)
	}
}
