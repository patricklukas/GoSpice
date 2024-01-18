package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {

	// bishopMoves := GenBishopLookups()
	// for i := 0; i < 64; i++ {
	// 	bishopMoves[i].Print()
	// }

	// Shift(FromSq(63), SE).Print()
	setupMasks()
	for i := 0; i < 64; i++ {
		kingMasks[i].Print()
		fmt.Println()
	}

	// for {
	// 	input := ""
	// 	fmt.Scanln(&input)
	// 	cmd := strings.Split(input, " ")[0]
	// 	switch cmd {
	// 	case "uci":
	// 		fmt.Println("uciok")
	// 	case "isready":
	// 		fmt.Println("readyok")
	// 	case "go":
	// 		fmt.Println("bestmove e7e5")
	// 	case "quit":
	// 		return
	// 	}
	// }
}