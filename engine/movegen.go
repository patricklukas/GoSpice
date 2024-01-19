package engine

// todo: generate enemy attacks
// todo: calculate checkmask
// todo: check attacks blocking castling

// generate all possible king moves
func GenKingMoves(brd Board) []Move {
	var moveList []Move
	// locate our king
	sq := lsb(brd.pieces[KING] & brd.friendly())
	// todo: find squares attacked by enemy pieces and exclude them
	moves := kingMasks[sq]
	// handle captures
	captures := moves & brd.enemy()
	for captures != 0 {
		// todo: avoid moving into check
		to := popLsb(&captures)
		moveList = append(moveList, NewMove(sq, to, KING, brd.pieceAt(to), EMPTY))
	}
	// handle quiet moves
	quiet := moves & brd.empty()
	for quiet != 0 {
		// todo: avoid moving into check
		to := popLsb(&quiet)
		moveList = append(moveList, NewMoveQuiet(sq, to, KING))
	}
	// todo: castling

	return moveList
}

func GenBishopMoves(brd Board) []Move {
	var moveList []Move
	// locate our bishops
	bishops := brd.pieces[BISHOP] & brd.friendly()
	// for each bishop
	for bishops != 0 {
		// todo: handle pins
		// todo: will use pinmasks and only allow bishop moves along the pin
		// get its square
		sq := popLsb(&bishops)
		// generate moves via raycasting
		moves := generateBishopAttacks(brd.occupied(), sq) & brd.checkmask
		// handle captures
		captures := moves & brd.enemy()
		for captures != 0 {
			to := popLsb(&captures)
			moveList = append(moveList, NewMove(sq, to, BISHOP, brd.pieceAt(to), EMPTY))
		}
		// handle quiet moves
		quiet := moves & brd.empty()
		for quiet != 0 {
			to := popLsb(&quiet)
			moveList = append(moveList, NewMoveQuiet(sq, to, BISHOP))
		}
	}
	return moveList
}
