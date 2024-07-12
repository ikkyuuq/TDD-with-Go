package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		got := Add(2, 3)
		expected := 5

		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	})
}

func ExampleAdd() {
	got := Add(1, 5)
	fmt.Println(got)
	// Output: 6
}
