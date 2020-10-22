package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("adding two integers", func(t *testing.T) {
		sum := Add(2,3)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("adding two integers", func(t *testing.T) {
		sum := Add(2,16)
		expected := 18

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("adding two integers", func(t *testing.T) {
		sum := Add(2,-16)
		expected := -14

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}