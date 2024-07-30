package chap02_integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected: %d, actual: %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
