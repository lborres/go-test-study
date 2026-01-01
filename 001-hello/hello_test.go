package hello

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("say Hello, World when providing empty string", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")

		assertCorrectMessage(t, got, want)
	})
	t.Run("say name when provided", func(t *testing.T) {
		want := "Hello, Kayne"
		got := Hello("Kayne", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola, Mira"
		got := Hello("Mira", "es")

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		want := "Bonjour, Aldo"
		got := Hello("Aldo", "fr")

		assertCorrectMessage(t, got, want)
	})
}
