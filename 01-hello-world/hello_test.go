package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		subject := "kennan"
		got := Hello(subject, "")
		want := "Hello, " + subject + "!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("just saying hello", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, there!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Bahasa Indonesia", func(t *testing.T) {
		subject := "kennan"
		got := Hello(subject, "Bahasa Indonesia")
		want := "Halo, " + subject + "!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		subject := "kennan"
		got := Hello(subject, "French")
		want := "Bonjour, " + subject + "!"

		assertCorrectMessage(t, got, want)
	})
}
