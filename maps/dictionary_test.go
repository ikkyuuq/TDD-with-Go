package dictionary

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorUnknownWord

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want.Error())
	})
}
