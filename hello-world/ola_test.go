package hello

import "testing"

func TestOla(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\n Got: %s \n Wanted: %s", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Ola("Albino")
		want := "Olá, Albino"

		verifyCorrectMessage(t, got, want)
	})

	t.Run("Say 'Olá, mundo!' when a empty string is passed", func(t *testing.T) {
		got := Ola("")
		want := "Olá, mundo!"
		verifyCorrectMessage(t, got, want)
	})
}
