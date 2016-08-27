package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "-"},
		[]string{"_", "_", "-"},
		[]string{"_", "_", "-"},
	}

	board[0][0] = "A"
	board[2][2] = "B"
	board[1][2] = "C"
	board[1][0] = "D"
	board[0][2] = "E"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
