package main

import "testing"

func TestHello(t *testing.T) {
	var got string
	var want string

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to country", func(t *testing.T) {
		got = Hello("Yiming", "cn")
		want = "Nihao, Yiming"
		assertCorrectMessage(t, got, want)

		got = Hello("Yiming", "es")
		want = "Hola, Yiming"
		assertCorrectMessage(t, got, want)

		got = Hello("Yiming", "fr")
		want = "Bonjour, Yiming"
		assertCorrectMessage(t, got, want)

		got = Hello("Yiming", "jp")
		want = "Soga, Yiming"
		assertCorrectMessage(t, got, want)

		got = Hello("Yiming", "en")
		want = "Hello, Yiming"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to world for empty string", func(t *testing.T) {
		got = Hello("Yiming", "")
		want = "Haha, world"
		assertCorrectMessage(t, got, want)
	})

}
