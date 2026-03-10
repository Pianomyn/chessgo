package main

import (
	"chessgo/board"
	"chessgo/movement"
	"fmt"
)

func main() {
	movement.PrintAllMoves(board.King)
	fmt.Println("Finished!")
}
