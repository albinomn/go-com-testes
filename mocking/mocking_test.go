package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("\n Got:  %q\n Want: %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Not enogh calls to sleeper\n Got: %d\n Want: 4", spySleeper.Calls)
	}
}
