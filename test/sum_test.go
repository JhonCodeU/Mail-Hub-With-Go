package main

import (
	"testing"
)

func Sum(a int, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(5, 5)
	}
}
