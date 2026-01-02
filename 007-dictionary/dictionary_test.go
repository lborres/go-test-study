package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "A procedure for critical evaluation"}

	assertDefinition := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	t.Run("should get definition of known term", func(t *testing.T) {
		term := "test"

		got, _ := dictionary.Search(term)
		want := "A procedure for critical evaluation"

		assertDefinition(t, got, want)

	})

	t.Run("should get error for unknown term", func(t *testing.T) {
		term := "unknown"
		_, err := dictionary.Search(term)

		want := "could not find definition for the given term"

		if err == nil {
			t.Error("expected to get an error.")
		}

		assertDefinition(t, err.Error(), want)

	})
}
