package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Ikkyuu", "")
		want := "Hello, Ikkyuu"
		assertMessage(t, got, want)
	})
	t.Run("Saying Hello, World! when an empty string is suppiled", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Beep", "French")
		want := "Bonjour, Beep"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
