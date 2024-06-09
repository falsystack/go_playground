package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dic := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dic.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q, given, %q", got, want, "test")
	}
}
