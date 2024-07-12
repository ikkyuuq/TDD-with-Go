package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("output of repeated", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa, 5 times"
		assert(t, repeated, expected)
	})
	t.Run("how many times of repeated", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb, 5 times"
		assert(t, repeated, expected)
	})
}

func assert(t testing.TB, repeated, expected string) {
	// t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
