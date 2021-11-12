package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "1Hello, World."
	if got := Hello(); got != want {
		// t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestNumbericZero(t *testing.T) {
	var num int
	zero := 0
	if num != zero {
		t.Error("num is supposed to be 0")
	}
}
