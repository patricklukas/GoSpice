package main

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

func (b BB) Pop() int {
	sq := int(BitScan(b))
	b.Unset(sq)
	return sq
}

// Index of least significant bit via debruijn algorithm
func BitScan(bb BB) int {
	const deBruijn64 uint64 = 0x03f79d71b4cb0a89
	deBruijnIndex := [64]int{0, 1, 48, 2, 57, 49, 28, 3, 61, 58, 50, 42, 38, 29, 17, 4, 62, 55, 59, 36, 53, 51, 43, 22, 45, 39, 33, 30, 24, 18, 12, 5, 63, 47, 56, 27, 60, 41, 37, 16, 54, 35, 52, 21, 44, 32, 23, 11, 46, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6}
	return deBruijnIndex[(uint64(bb&-bb)*deBruijn64)>>58]
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
