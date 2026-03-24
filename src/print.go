// https://www.chessprogramming.org/Knight_Pattern
package main

import (
	"chessgo/board"
	"chessgo/movement"
	"fmt"
)

func PrintMoves(sourceSquare board.Square, bb board.Bitboard) {
	fmt.Println("  +-----------------+")
	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1) // Rank number
		for f := 0; f < 8; f++ {
			sq := board.Square(r*8 + f)

			if sq == sourceSquare {
				fmt.Print("O ")
				continue
			}

			if (uint64(bb)>>uint8(sq))&1 == 1 {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
}

func PrintAllMoves(piece board.Piece) {
	var allMoves []board.Bitboard
	switch piece {
	case board.King:
		allMoves = movement.GetKingAttacks()
	case board.Knight:
		allMoves = movement.GetKnightAttacks()
	case board.Rook:
		allMoves = movement.GetRookAttacks()
	case board.Bishop:
		allMoves = movement.GetBishopAttacks()
	case board.Queen:
		allMoves = movement.GetQueenAttacks()
	default:
		panic("Invalid Piece")
	}
	for i, bb := range allMoves {
		PrintMoves(board.Square(i), bb)
	}
}
