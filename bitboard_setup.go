// Some ideas taken and adapted from https://github.com/stephenjlovell/gopher_check/blob/master/bitboard_setup.go
// Thank you Stephen Lovell!

package main

var pawnOffsets = [2][2]int{{8, 16}, {-8, -16}}
var pawnAttackOffsets = [2][2]int{{7, 9}, {-7, -9}}
var knightOffsets = [8]int{-17, -15, -10, -6, 6, 10, 15, 17}
var bishopDirs = [4]int{NE, SE, SW, NW}
var rookDirs = [4]int{NO, EA, SO, WE}

var dirs = [8]int{8, 9, 1, -7, -8, -9, -1, 7}

var rowMask, colMask [8]BB
var rayMasks [64][8]BB
var sqMaskOn, sqMaskOff [64]BB
var knightMasks, bishopMasks, rookMasks, queenMasks, kingMasks [64]BB
var pawnMasks, pawnAttackMasks [2][64]BB

func manhattanDistance(from, to int) int {
	return abs(RankIdx(from)-RankIdx(to)) + abs(FileIdx(from)-FileIdx(to))
}

func setupRowAndColMasks() {
	for i := 0; i < 8; i++ {
		rowMask[i] = 0xff << (i * 8)
		colMask[i] = 0x0101010101010101 << i
	}
}

func setupSquareMasks() {
	for i := 0; i < 64; i++ {
		sqMaskOn[i] = BB(1 << i)
		sqMaskOff[i] = ^sqMaskOn[i]
	}
}

func setupKnightMasks() {
	for i := 0; i < 64; i++ {
		for _, offset := range knightOffsets {
			sq := i + offset
			if onBoard(sq) && manhattanDistance(i, sq) == 3 {
				knightMasks[i] |= sqMaskOn[sq]
			}
		}
	}
}

func setupBishopMasks() {
	for i := 0; i < 64; i++ {
		for _, dir := range bishopDirs {
			rayMasks[i][dir] = 0
			offset := dirs[dir]
			last := i
			for sq := i + offset; onBoard(sq) && manhattanDistance(last, sq) == 2; sq += offset {
				rayMasks[i][dir] |= sqMaskOn[sq]
				last = sq
			}
			bishopMasks[i] |= rayMasks[i][dir]
		}
	}
}

func setupRookMasks() {
	for i := 0; i < 64; i++ {
		for _, dir := range rookDirs {
			rayMasks[i][dir] = 0
			offset := dirs[dir]
			last := i
			for sq := i + offset; onBoard(sq) && manhattanDistance(last, sq) == 1; sq += offset {
				rayMasks[i][dir] |= sqMaskOn[sq]
				last = sq
			}
			rookMasks[i] |= rayMasks[i][dir]
		}
	}
}

func setupQueenMasks() {
	for i := 0; i < 64; i++ {
		queenMasks[i] = bishopMasks[i] | rookMasks[i]
	}
}

func setupKingMasks() {
	for i := 0; i < 64; i++ {
		for _, dir := range dirs {
			sq := i + dir
			if onBoard(sq) && manhattanDistance(i, sq) <= 2 {
				kingMasks[i] |= sqMaskOn[sq]
			}
		}
	}
}

func setupPawnMasks() {
	for side := 0; side < 2; side++ {
		// exclude first and last rank as no pawns can exist there
		for i := 8; i < 56; i++ {
			sq := i + pawnOffsets[side][0]
			if onBoard(sq) && manhattanDistance(i, sq) == 1 {
				pawnMasks[side][i] |= sqMaskOn[sq]
			}
			// double push on starting rank
			if FromSq(i)&rowMask[1+5*side] != 0 {
				sq = i + pawnOffsets[side][1]
				pawnMasks[side][i] |= sqMaskOn[sq]
			}
			// attacks
			for _, offset := range pawnAttackOffsets[side] {
				sq = i + offset
				if onBoard(sq) && manhattanDistance(i, sq) == 2 {
					pawnAttackMasks[side][i] |= sqMaskOn[sq]
				}
			}
		}
	}
}

func setupMasks() {
	setupRowAndColMasks()
	setupSquareMasks()
	setupKnightMasks()
	setupBishopMasks()
	setupRookMasks()
	setupQueenMasks()
	setupKingMasks()
	setupPawnMasks()
}
