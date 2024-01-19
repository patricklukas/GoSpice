package engine

// todo: does move need expansion?
// todo: -> ep, castle
// todo: alternative in visitor pattern style?

func genEnemyAttacks(brd Board) {
	attacks := BB(0)
	checkmask := BB(0)
	// locate enemy pieces
	enemy := brd.enemy()
	friendlyKing := brd.pieces[KING] & brd.friendly()
	friendlyKingSq := lsb(friendlyKing)
	// for each piece
	for enemy != 0 {
		// get its square
		sq := popLsb(&enemy)
		// generate attacks
		switch brd.pieceAt(sq) {
		case PAWN:
			attacks |= pawnAttackMasks[flipColor(brd.turn)][sq]
			if attacks&friendlyKing != 0 {
				checkmask |= sqMaskOn[sq]
			}
		case KNIGHT:
			attacks |= knightMasks[sq]
			if attacks&friendlyKing != 0 {
				checkmask |= sqMaskOn[sq]
			}
		case BISHOP:
			attacks |= generateBishopAttacks(brd.occupied(), sq)
			if attacks&friendlyKing != 0 {
				// add only line from king to attacking bishop to checkmask
				checkmask |= generateKingBishopRay(sqMaskOn[sq], friendlyKingSq)
			}
		case ROOK:
			attacks |= generateRookAttacks(brd.occupied(), sq)
			if attacks&friendlyKing != 0 {
				// add only line from king to attacking rook to checkmask
				checkmask |= generateKingRookRay(sqMaskOn[sq], friendlyKingSq)
			}
		case QUEEN:
			attacks |= generateQueenAttacks(brd.occupied(), sq)
			if attacks&friendlyKing != 0 {
				// add only line from king to attacking queen to checkmask
				checkmask |= generateKingQueenRay(sqMaskOn[sq], friendlyKingSq)
			}
		case KING:
			attacks |= kingMasks[sq]
			// king cannot attack other king
		}
	}
	brd.enemyAttacks = attacks
	if checkmask > 0 {
		brd.checkmask = checkmask
	} else {
		brd.resetCheckmask()
	}
}

func genPinmasks(brd Board) {
	friendlyKingSq := lsb(brd.pieces[KING] & brd.friendly())
	enemyQueensRooks := brd.enemy() & (brd.pieces[QUEEN] | brd.pieces[ROOK])
	enemyQueensBishops := brd.enemy() & (brd.pieces[QUEEN] | brd.pieces[BISHOP])
	pinnables := brd.occupied() & ^(enemyQueensBishops | enemyQueensRooks)
	// find xrays
	brd.hvPins = genPinsHV(pinnables, enemyQueensRooks, friendlyKingSq)
	brd.diagPins = genPinsDiag(pinnables, enemyQueensBishops, friendlyKingSq)
}

// generate all possible king moves
func genKingMoves(brd Board) []Move {
	var moveList []Move
	// locate our king
	sq := lsb(brd.pieces[KING] & brd.friendly())
	// lookup king moves and exclude enemy attacks which also
	// avoids moving into check
	moves := kingMasks[sq] & ^brd.enemyAttacks
	// handle captures
	captures := moves & brd.enemy()
	for captures != 0 {
		to := popLsb(&captures)
		moveList = append(moveList, NewMove(sq, to, KING, brd.pieceAt(to), EMPTY))
	}
	// handle quiet moves
	quiet := moves & brd.empty()
	for quiet != 0 {
		to := popLsb(&quiet)
		moveList = append(moveList, NewMoveQuiet(sq, to, KING))
	}
	// castling
	if brd.castleL[brd.turn] {
		if castlingMasks[brd.turn][0]&brd.enemyAttacks|kingToRookMasks[brd.turn][0]&brd.empty() == 0 {
			moveList = append(moveList, NewMoveQuiet(4+int(56*brd.turn), 2+int(56*brd.turn), KING))
		}
	}
	if brd.castleR[brd.turn] {
		if castlingMasks[brd.turn][1]&brd.enemyAttacks|kingToRookMasks[brd.turn][1]&brd.empty() == 0 {
			moveList = append(moveList, NewMoveQuiet(4+int(56*brd.turn), 6+int(56*brd.turn), KING))
		}
	}
	// check if squares between king and rook are empty

	// castles := brd.castlingRights(brd.turn)

	return moveList
}

func genPawnMoves(brd Board) []Move {
	var moveList []Move
	// locate our pawns
	pawns := brd.pieces[PAWN] & brd.friendly()
	// for each pawn
	for pawns != 0 {
		// get its square
		sq := popLsb(&pawns)
		// generate captures
		captures := pawnAttackMasks[brd.turn][sq] & (brd.enemy() | (brd.ep & ^brd.diagPins))
		// if a pawn is member of the diagonal pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.diagPins != 0 {
			captures &= brd.diagPins
		}
		for captures != 0 {
			to := popLsb(&captures)
			moveList = append(moveList, NewMove(sq, to, PAWN, brd.pieceAt(to), EMPTY))
		}
		// generate moves
		quiet := pawnMasks[brd.turn][sq] & brd.checkmask & brd.empty()
		// if a pawn is member of the hv pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.hvPins != 0 {
			quiet &= brd.hvPins
		}
		for quiet != 0 {
			to := popLsb(&quiet)
			moveList = append(moveList, NewMoveQuiet(sq, to, PAWN))
		}
	}
	return moveList
}

func genKnightMoves(brd Board) []Move {
	var moveList []Move
	// locate our knights
	knights := brd.pieces[KNIGHT] & brd.friendly() & ^(brd.hvPins | brd.diagPins)
	// for each knight
	for knights != 0 {
		// get its square
		sq := popLsb(&knights)
		// generate moves
		moves := knightMasks[sq] & brd.checkmask
		// handle captures
		captures := moves & brd.enemy()
		for captures != 0 {
			to := popLsb(&captures)
			moveList = append(moveList, NewMove(sq, to, KNIGHT, brd.pieceAt(to), EMPTY))
		}
		// handle quiet moves
		quiet := moves & brd.empty()
		for quiet != 0 {
			to := popLsb(&quiet)
			moveList = append(moveList, NewMoveQuiet(sq, to, KNIGHT))
		}
	}
	return moveList
}

func genQueenMoves(brd Board) []Move {
	var moveList []Move
	// locate our queens
	queens := brd.pieces[QUEEN] & brd.friendly()
	// for each queen
	for queens != 0 {
		// get its square
		sq := popLsb(&queens)
		// generate moves via raycasting
		moves := generateQueenAttacks(brd.occupied(), sq) & brd.checkmask
		// if a queen is member of the diagonal pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.diagPins != 0 {
			moves &= brd.diagPins
		}
		// if a queen is member of the hv pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.hvPins != 0 {
			moves &= brd.hvPins
		}
		// handle captures
		captures := moves & brd.enemy()
		for captures != 0 {
			to := popLsb(&captures)
			moveList = append(moveList, NewMove(sq, to, QUEEN, brd.pieceAt(to), EMPTY))
		}
		// handle quiet moves
		quiet := moves & brd.empty()
		for quiet != 0 {
			to := popLsb(&quiet)
			moveList = append(moveList, NewMoveQuiet(sq, to, QUEEN))
		}
	}
	return moveList
}

func genBishopMoves(brd Board) []Move {
	var moveList []Move
	// locate our bishops
	bishops := brd.pieces[BISHOP] & brd.friendly()
	// for each bishop
	for bishops != 0 {
		// get its square
		sq := popLsb(&bishops)
		// generate moves via raycasting
		moves := generateBishopAttacks(brd.occupied(), sq) & brd.checkmask
		// if a bishop is member of the diagonal pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.diagPins != 0 {
			moves &= brd.diagPins
		}
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

func genRookMoves(brd Board) []Move {
	var moveList []Move
	// locate our rooks
	rooks := brd.pieces[ROOK] & brd.friendly()
	// for each rook
	for rooks != 0 {
		// get its square
		sq := popLsb(&rooks)
		// generate moves via raycasting
		moves := generateRookAttacks(brd.occupied(), sq) & brd.checkmask
		// if a rook is member of the hv pinmask, it can only move along the pin
		if sqMaskOn[sq]&brd.hvPins != 0 {
			moves &= brd.hvPins
		}
		// handle captures
		captures := moves & brd.enemy()
		for captures != 0 {
			to := popLsb(&captures)
			moveList = append(moveList, NewMove(sq, to, ROOK, brd.pieceAt(to), EMPTY))
		}
		// handle quiet moves
		quiet := moves & brd.empty()
		for quiet != 0 {
			to := popLsb(&quiet)
			moveList = append(moveList, NewMoveQuiet(sq, to, ROOK))
		}
	}
	return moveList
}
