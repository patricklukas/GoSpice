package main

import (
	"testing"
)

func BenchmarkSetupMasks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setupMasks()
	}
}
