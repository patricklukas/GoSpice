package engine

import "testing"

func TestMoveFrom(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if m.From() != 1 {
		t.Errorf("Expected From() to be 1, got %d", m.From())
	}
}

func TestMoveTo(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if m.To() != 2 {
		t.Errorf("Expected To() to be 2, got %d", m.To())
	}
}

func TestMovePiece(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if m.Piece() != 3 {
		t.Errorf("Expected Piece() to be 3, got %d", m.Piece())
	}
}

func TestMoveCapturedPiece(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if m.CapturedPiece() != 4 {
		t.Errorf("Expected CapturedPiece() to be 4, got %d", m.CapturedPiece())
	}
}

func TestMoveIsCapture(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if !m.isCapture() {
		t.Errorf("Expected isCapture() to be true, got false")
	}
}

func TestMoveIsQuiet(t *testing.T) {
	m := NewMove(1, 2, 3, EMPTY, EMPTY)
	if !m.isQuiet() {
		t.Errorf("Expected isQuiet() to be true, got false")
	}
}

func TestNewMove(t *testing.T) {
	m := NewMove(1, 2, 3, 4, 5)
	if m.From() != 1 || m.To() != 2 || m.Piece() != 3 || m.CapturedPiece() != 4 {
		t.Errorf("NewMove() didn't correctly encode the move")
	}
}

func TestMovePromoteToQueen(t *testing.T) {
	m := NewMove(1, 2, 3, 4, QUEEN)
	if m.PromoteTo() != QUEEN {
		t.Errorf("Expected PromoteToPiece() to be 4, got %d", m.PromoteTo())
	}
}

func TestMovePromoteToKnight(t *testing.T) {
	m := NewMove(1, 2, 3, 4, KNIGHT)
	if m.PromoteTo() != KNIGHT {
		t.Errorf("Expected PromoteToPiece() to be 1, got %d", m.PromoteTo())
	}
}
