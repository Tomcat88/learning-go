package main

import (
	"bytes"
	"testing"
)

func TestInjection(t *testing.T) {
	t.Run("greet should greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Thomas")

		got := buffer.String()
		want := "Hello, Thomas"
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
