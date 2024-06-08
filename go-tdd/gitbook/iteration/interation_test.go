package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got %q, expected %q", repeated, expected)
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := Repeat("b", 10)
	expected := "bbbbbbbbbb"

	if repeated != expected {
		t.Errorf("Got %q, expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
