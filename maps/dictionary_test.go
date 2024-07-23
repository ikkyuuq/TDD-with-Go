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
		want := ErrNotFound

		// Expected error
		assertError(t, err, ErrNotFound)

		assertStrings(t, err.Error(), want.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		definition := "this is just a test"
		dictionary.Add(key, definition)

		assertDefinition(t, dictionary, key, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		err := dictionary.Add(key, "new test")

		// Expected error
		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, key, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertError(t testing.TB, err error, expected DictionaryErr) {
	t.Helper()
	if err.Error() != expected.Error() {
		t.Fatal("expected to get an error")
	}
}

func TestUpdate(t *testing.T) {
	key := "test"
	definition := "this is just a test"
	dictionary := Dictionary{key: definition}
	newDefinition := "new definition"

	dictionary.Update(key, newDefinition)

	assertDefinition(t, dictionary, key, newDefinition)
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	assertError(t, err, ErrNotFound)
}
