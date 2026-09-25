package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Greg")
	want := "Hello, Greg"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
