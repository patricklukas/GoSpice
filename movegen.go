package main

// this is a pretty ugly way to do this, but let's hope it works and think about refactoring later
func GenPawnMoves(brd Board, clrToMove int) []Move {
	moves := []Move{}
	// is there a cleaner way to do this?
	flip := clrToMove*-2 + 1

	// generate single push moves
	singlePushes := brd.pieces[PAWN] & brd.colors[clrToMove]
	singlePushes = Shift(singlePushes, flip*NO) & brd.Empty()
	// use single pushes to avoid double checking for empty squares
	doublePushes := Shift(singlePushes&rowMask[2], flip*NO) & brd.Empty()

	// iterate over single pushes until empty...
	for singlePushes != 0 {
		// ...pop them
		to := singlePushes.Pop()
		// ...store the from square
		from := to - -flip*SO
		// ...handle promotions
		if to >= 48 {
			// let's ignore rook and bishop promotions as they are always inferior to queen promotion
			moves = append(moves, NewMove(from, to, PAWN, EMPTY, QUEEN))
			moves = append(moves, NewMove(from, to, PAWN, EMPTY, KNIGHT))
		} else {
			// ...and add the move to the list
			moves = append(moves, NewMove(from, to, PAWN, EMPTY, EMPTY))
		}
	}

	// now the same for double pushes
	for doublePushes != 0 {
		to := doublePushes.Pop()
		// we should be protected from udnerflow here, as moves are only generated from row 2
		from := to - -2*flip*SO
		moves = append(moves, NewMove(from, to, PAWN, EMPTY, EMPTY))
	}

	// generate captures
	captures := brd.pieces[PAWN] & brd.colors[clrToMove]
	capturesNW := Shift(captures, flip*NW) & brd.colors[clrToMove^1]
	capturesNE := Shift(captures, flip*NE) & brd.colors[clrToMove^1]

	// iterate over captures until empty...
	for capturesNW != 0 {
		// ...pop them
		to := capturesNW.Pop()
		// ...store the from square
		from := to - -flip*SE

		capturedPiece := EMPTY
		// find captured piece
		for i := 0; i <= int(QUEEN); i++ {
			if brd.pieces[i]&brd.colors[clrToMove^1]&BB(to) != 0 {
				capturedPiece = Piece(i)
			}
		}

		// ...handle promotions
		if to >= 48 {
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, QUEEN))
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, KNIGHT))
		} else {
			// ...and add the move to the list
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, EMPTY))
		}
	}
	// iterate over captures until empty...
	for capturesNE != 0 {
		// ...pop them
		to := capturesNE.Pop()
		// ...store the from square
		from := to - -flip*SW

		capturedPiece := EMPTY
		// find captured piece
		for i := 0; i <= int(QUEEN); i++ {
			if brd.pieces[i]&brd.colors[clrToMove^1]&BB(to) != 0 {
				capturedPiece = Piece(i)
			}
		}

		// ...handle promotions
		if to >= 48 {
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, QUEEN))
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, KNIGHT))
		} else {
			// ...and add the move to the list
			moves = append(moves, NewMove(from, to, PAWN, capturedPiece, EMPTY))
		}
	}

	// handle en passant captures
	if brd.ep != 0 {
		capturesEpNW := Shift(captures, flip*NW) & FromSq(int(brd.ep))
		capturesEpNE := Shift(captures, flip*NE) & FromSq(int(brd.ep))

		for capturesEpNW != 0 {
			to := capturesEpNW.Pop()
			from := to - -flip*SE
			moves = append(moves, NewMove(from, to, PAWN, PAWN, EMPTY))
		}
		for capturesEpNE != 0 {
			to := capturesEpNE.Pop()
			from := to - -flip*SW
			moves = append(moves, NewMove(from, to, PAWN, PAWN, EMPTY))
		}
	}

	return moves
}

func GenKingMoves(brd Board, clrToMove int) []Move {
	moves := []Move{}
	// is there a cleaner way to do this?
	flip := clrToMove*-2 + 1

	// generate king moves
	king := brd.pieces[KING] & brd.colors[clrToMove]
	kingMoves := Shift(king, NO) | Shift(king, SO) | Shift(king, EA) | Shift(king, WE) | Shift(king, NE) | Shift(king, NW) | Shift(king, SE) | Shift(king, SW)
	kingMoves &= ^brd.colors[clrToMove]

	// iterate over single pushes until empty...
	for kingMoves != 0 {
		// ...pop them
		to := kingMoves.Pop()
		// ...store the from square
		from := to - -flip*SO

		// ...find captured piece
		capturedPiece := EMPTY
		for i := 0; i <= int(QUEEN); i++ {
			if brd.pieces[i]&brd.colors[clrToMove^1]&BB(to) != 0 {
				capturedPiece = Piece(i)
			}
		}

		// ...and add the move to the list
		moves = append(moves, NewMove(from, to, KING, capturedPiece, EMPTY))
	}

	return moves
}

func GenKnightMoves(brd Board, clrToMove int) []Move {
	moves := []Move{}
	// is there a cleaner way to do this?
	flip := clrToMove*-2 + 1

	// generate knight moves
	knight := brd.pieces[KNIGHT] & brd.colors[clrToMove]
	knightMoves := Shift(knight, NO+NO+EA) | Shift(knight, NO+NO+WE) | Shift(knight, SO+SO+EA) | Shift(knight, SO+SO+WE) | Shift(knight, EA+EA+NO) | Shift(knight, EA+EA+SO) | Shift(knight, WE+WE+NO) | Shift(knight, WE+WE+SO)
	knightMoves &= ^brd.colors[clrToMove]

	// iterate over single pushes until empty...
	for knightMoves != 0 {
		// ...pop them
		to := knightMoves.Pop()
		// ...store the from square
		from := to - -flip*SO

		// ...find captured piece
		capturedPiece := EMPTY
		for i := 0; i <= int(QUEEN); i++ {
			if brd.pieces[i]&brd.colors[clrToMove^1]&BB(to) != 0 {
				capturedPiece = Piece(i)
			}
		}

		// ...and add the move to the list
		moves = append(moves, NewMove(from, to, KNIGHT, capturedPiece, EMPTY))
	}

	return moves
}

func GenBishopMoves(brd Board, clrToMove int) []Move {
	moves := []Move{}
	// is there a cleaner way to do this?
	flip := clrToMove*-2 + 1

	// generate bishop moves
	bishop := brd.pieces[BISHOP] & brd.colors[clrToMove]
	bishopMoves := bishop
	for i := 0; i < 7; i++ {
		bishopMoves |= Shift(bishopMoves, NO)
		bishopMoves &= ^rowMask[7]
	}
	bishopMoves &= ^brd.colors[clrToMove]

	// iterate over single pushes until empty...
	for bishopMoves != 0 {
		// ...pop them
		to := bishopMoves.Pop()
		// ...store the from square
		from := to - -flip*SO

		// ...find captured piece
		capturedPiece := EMPTY
		for i := 0; i <= int(QUEEN); i++ {
			if brd.pieces[i]&brd.colors[clrToMove^1]&BB(to) != 0 {
				capturedPiece = Piece(i)
			}
		}

		// ...and add the move to the list
		moves = append(moves, NewMove(from, to, BISHOP, capturedPiece, EMPTY))
	}

	return moves
}

// Generate all possible bishop moves for a given square
func GenBishopLookups() [64]BB {
	var lookups [64]BB

	for sq := 0; sq < 64; sq++ {
		start := FromSq(sq)
		bishop := start
		for i := 0; i < 7; i++ {
			bishop |= Shift(start, i*NE)
			bishop |= Shift(start, i*SE)
			bishop |= Shift(start, i*NW)
			bishop |= Shift(start, i*SW)
		}
		lookups[sq] = bishop
	}

	return lookups
}
