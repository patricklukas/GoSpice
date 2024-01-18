package main

func main() {
	InitBoards()
	BB.Print(brd.pieces[PAWN] & brd.pieces[WHITE])
}
