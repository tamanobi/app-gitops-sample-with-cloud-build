package main

import "testing"

func TestHelloReturnsHello(t *testing.T) {
	got := Hello()
	want := "Hello"
	if got != want {
		t.Fatalf("Hello() = %v, but got %v\n", want, got)
	}
}
