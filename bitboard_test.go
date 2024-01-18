package main

import "testing"

var testBB = FromSq(A2)

func BenchmarkMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBB.Move(A2, A3)
	}
}

func BenchmarkMoveByInversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBB.MoveByInversion(A2, A3)
	}
}
