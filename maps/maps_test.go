package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("Unkown word", func(t *testing.T) {
		_, err := dictionary.Search("tester")

		if err == nil {
			t.Fatal("Expected to get errors")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\n Got:   %q\n Want:  %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("\n Got:   %q\n Want:  %q", got, want)
	}
}
