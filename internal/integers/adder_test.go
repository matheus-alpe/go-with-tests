package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add should return 4", func(t *testing.T) {
		assertInteger(t, 4, Add(2, 2))
	})

	t.Run("add should return 5 on x=2 and y=3", func(t *testing.T) {
		assertInteger(t, 5, Add(2, 3))
	})
}

func assertInteger(t testing.TB, expected, want int) {
	t.Helper()

	if expected != want {
		t.Errorf("expected %d but got %d", expected, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
