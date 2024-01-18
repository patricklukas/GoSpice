package main

import (
	"testing"
)

func TestSqToCoord(t *testing.T) {
	tests := []struct {
		sq    int
		coord string
	}{
		{A1, "A1"},
		{B1, "B1"},
		{C1, "C1"},
		// Add more test cases here
	}

	for _, tt := range tests {
		result := SqToCoord(tt.sq)
		if result != tt.coord {
			t.Errorf("SqToCoord(%d) = %s; want %s", tt.sq, result, tt.coord)
		}
	}
}

func TestCoordToSq(t *testing.T) {
	tests := []struct {
		coord string
		sq    int
	}{
		{"A1", A1},
		{"B1", B1},
		{"C1", C1},
		// Add more test cases here
	}

	for _, tt := range tests {
		result := CoordToSq(tt.coord)
		if result != tt.sq {
			t.Errorf("CoordToSq(%s) = %d; want %d", tt.coord, result, tt.sq)
		}
	}
}
