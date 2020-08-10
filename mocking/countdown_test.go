package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	// want := "3\n"
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

	if spySleeper.NumCalls != 4 {
		t.Errorf("expected number of calls: 4, got: %d", spySleeper.NumCalls)
	}
}
