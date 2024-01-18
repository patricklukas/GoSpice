package main

type Move uint32

// Encode a move into 18 bits according to this scheme:
// Least Significant Bit First
// b0-5: from square
// b6-11: to square
// b12-14: piece
// b15-17: captured piece

// extract lowest 6 bits encoding the from square
func (m Move) From() int {
	return int(m & 63)
}

// extract bits 6 to 11, encoding the to square
func (m Move) To() int {
	return int(m >> 6 & 63)
}

func (m Move) Piece() int {
	return int(m >> 12 & 7)
}

func (m Move) CapturedPiece() Piece {
	return Piece(m >> 15 & 7)
}

func (m Move) isCapture() bool {
	return m.CapturedPiece() != EMPTY
}

func (m Move) isQuiet() bool {
	return !m.isCapture()
}

func NewMove(from, to int, piece, capturedPiece Piece) Move {
	return Move(from) | Move(to)<<6 | Move(piece)<<12 | Move(capturedPiece)<<15
}
