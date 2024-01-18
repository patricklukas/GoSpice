package engine

import (
	"fmt"
	"strings"
)

func HandleUci() {
	for {
		input := ""
		fmt.Scanln(&input)
		cmd := strings.Split(input, " ")
		switch cmd[0] {
		case "uci":
			fmt.Println("id name engine 0.1")
			fmt.Println("id author Patrick Hein")
			fmt.Println("uciok")
		case "isready":
			fmt.Println("readyok")
		case "position":
			if cmd[1] == "startpos" {

			} else {

			}
		case "go":
			fmt.Println("bestmove e7e5")
		case "quit":
			return
		}
	}
}
