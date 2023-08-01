package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	r := Repeat("a", 5)
	expected := "aaaaa"

	if r != expected {
		t.Errorf("expected %q but found %q", expected, r)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	r := Repeat("hello", 3)
	fmt.Println(r)
	// Output: hellohellohello
}
