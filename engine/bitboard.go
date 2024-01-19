package engine

import (
	"fmt"
)

// Wrapper for uint64 representing a bitboard
type BB uint64

func FromSq(sq int) BB {
	return BB(1 << sq)
}

// Get a square from bitboard
func (b BB) Get(sq int) BB {
	return b & FromSq(sq)
}

// Set a square on bitboard
func (b *BB) Set(sq int) {
	*b |= FromSq(sq)
}

// Unset a square on bitboard
func (b *BB) Unset(sq int) {
	bb := FromSq(sq)
	*b |= bb
	*b ^= bb
}

func (b *BB) ShiftInplace(n int) {
	if n >= 0 {
		*b <<= uint(n)
	} else {
		*b >>= uint(-n)
	}
}

func Shift(b BB, n int) BB {
	if n >= 0 {
		return b << uint(n)
	} else {
		return b >> uint(-n)
	}
}

// move a bit from square to square
func (b *BB) Move(from int, to int) {
	b.Unset(from)
	b.Set(to)
}

// move a bit from square to square
func (b *BB) MoveByInversion(from int, to int) {
	*b ^= FromSq(from) | FromSq(to)
}

func (b BB) IsSet(sq int) bool {
	return b&FromSq(sq) != 0
}

func (b BB) IsEmpty() bool {
	return b == 0
}

// Print bit representation of uint 64 as 8 by 8 board
func (b BB) Print() {
	for i := uint64(7); i < 8; i-- {
		// fmt.Printf("%d ", i+1)
		for a := uint64(0); a < 8; a++ {
			fmt.Printf("%d ", 1&(b>>(a+i*8)))
		}
		fmt.Println()
	}
	// fmt.Println("  A B C D E F G H")
}
