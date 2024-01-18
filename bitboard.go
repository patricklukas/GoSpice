package main

import "fmt"

// Wrapper for uint64 representing a bitboard
type BB uint64

func FromSq(sq Sq) BB {
	return BB(1 << sq)
}

// Set a square on bitboard
func (b *BB) Set(sq Sq) {
	*b |= FromSq(sq)
}

// Clear a square on bitboard
func (b *BB) Clear(sq Sq) {
	bb := FromSq(sq)
	*b |= bb
	*b ^= bb
}

// move a bit from square to square
func (b *BB) Move(from Sq, to Sq) {
	b.Clear(from)
	b.Set(to)
}

// move a bit from square to square
func (b *BB) MoveByInversion(from Sq, to Sq) {
	*b ^= FromSq(from) | FromSq(to)
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
