package engine

// Holding all relevant information about the board state
type Board struct {
	pieces    [6]BB //384 bits
	colors    [2]BB //128 bits
	hvPins    [2]BB // 128 bits
	diagPins  [2]BB // 128 bits
	checkmask BB    // 64 bits
	ep        BB    // 64 bits
	castling  uint8 // 8 bits
	turn      uint8 // 8 bits
}

var brd Board

func (brd *Board) occupied() BB {
	return brd.colors[WHITE] | brd.colors[BLACK]
}

func (brd *Board) empty() BB {
	return ^brd.occupied()
}

func (brd *Board) friendly() BB {
	return brd.colors[brd.turn]
}

func (brd *Board) enemy() BB {
	return brd.colors[flipColor(brd.turn)]
}

func (brd *Board) enemyOrEmpty() BB {
	return ^brd.colors[brd.turn]
}

func (brd *Board) pieceAt(sq int) Piece {
	for piece, bb := range brd.pieces {
		if bb.Get(sq) != 0 {
			return Piece(piece)
		}
	}
	return EMPTY
}

func (brd *Board) resetCheckmask() {
	brd.checkmask = 0xFFFFFFFFFFFFFFFF
}

func InitBoards() {
	brd.colors[WHITE] = rowMask[0] | rowMask[1]
	brd.colors[BLACK] = rowMask[6] | rowMask[7]
	brd.pieces[PAWN] = rowMask[1] | rowMask[6]
	brd.pieces[ROOK].Set(A1)
	brd.pieces[ROOK].Set(H1)
	brd.pieces[ROOK].Set(A8)
	brd.pieces[ROOK].Set(H8)
	brd.pieces[KNIGHT].Set(B1)
	brd.pieces[KNIGHT].Set(G1)
	brd.pieces[KNIGHT].Set(B8)
	brd.pieces[KNIGHT].Set(G8)
	brd.pieces[BISHOP].Set(C1)
	brd.pieces[BISHOP].Set(F1)
	brd.pieces[BISHOP].Set(C8)
	brd.pieces[BISHOP].Set(F8)
	brd.pieces[QUEEN].Set(D1)
	brd.pieces[QUEEN].Set(D8)
	brd.pieces[KING].Set(E1)
	brd.pieces[KING].Set(E8)
	brd.resetCheckmask()
	brd.turn = WHITE
	brd.castling = 0xF
}
