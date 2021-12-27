package integers

import (
	"fmt"
	"testing"
)

func ExampleAdder() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", sum, expected)
	}
}
