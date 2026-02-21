// https://www.chessprogramming.org/Knight_Pattern
package movement

import (
	"chessgo/board"
	"fmt"
)

const (
	NoWeWe int8 = 6
	NoNoWe int8 = 15
	NoEaEa int8 = 10
	NoNoEa int8 = 17
	SoWeWe int8 = -10
	SoSoWe int8 = -17
	SoEaEa int8 = -6
	SoSoEa int8 = -15
)

func CreateKnightAttackTable() []board.Bitboard {
	knightAttacks := make([]board.Bitboard, 64)
	for i := range knightAttacks {
		bit := board.Bitboard(1) << i

		knightAttacks[i] |= (bit & board.NotFileABMask) << NoWeWe
		knightAttacks[i] |= (bit & board.NotFileAMask) << NoNoWe
		knightAttacks[i] |= (bit & board.NotFileGHMask) << NoEaEa
		knightAttacks[i] |= (bit & board.NotFileHMask) << NoNoEa

		// Cant negative bitshift
		knightAttacks[i] |= (bit & board.NotFileABMask) >> (-SoWeWe)
		knightAttacks[i] |= (bit & board.NotFileAMask) >> (-SoSoWe)
		knightAttacks[i] |= (bit & board.NotFileGHMask) >> (-SoEaEa)
		knightAttacks[i] |= (bit & board.NotFileHMask) >> (-SoSoEa)
	}

	return knightAttacks
}

func PrintAllKnightMoves() {
	bb := CreateKnightAttackTable()
	for i, square := range bb {
		PrintKnightMoves(board.Square(i), square)
	}
}

func PrintKnightMoves(sourceSquare board.Square, bb board.Bitboard) {
	fmt.Println("  +-----------------+")
	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1) // Rank number
		for f := 0; f < 8; f++ {
			sq := board.Square(r*8 + f)

			if sq == sourceSquare {
				fmt.Print("N ")
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
