package main

import "testing"

func BenchmarkLoopI(b *testing.B) {

	s := make([]struct{}, b.N)

	for i := 0; i < len(s); i++ {
	}
}

func BenchmarkLoopRange(b *testing.B) {
	s := make([]struct{}, b.N)
	for _, _ = range s {
	}
}
