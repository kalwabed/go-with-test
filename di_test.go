package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ali")

	got := buffer.String()
	want := "Hello, Ali"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
