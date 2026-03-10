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

const (
	N  int8 = 8
	NE int8 = 9
	E  int8 = 1
	SE int8 = -7
	S  int8 = -8
	SW int8 = -9
	W  int8 = -1
	NW int8 = 7
)

func CreateKingAttackTable() []board.Bitboard {
	kingAttacks := make([]board.Bitboard, 64)
	for i := range kingAttacks {
		bit := board.Bitboard(1) << i

		kingAttacks[i] |= (bit & board.NotRank8Mask) << (N)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank8Mask) << (NE)
		kingAttacks[i] |= (bit & board.NotFileHMask) << (E)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank1Mask) >> (-SE)
		kingAttacks[i] |= (bit & board.NotRank1Mask) >> (-S)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank1Mask) >> (-SW)
		kingAttacks[i] |= (bit & board.NotFileAMask) >> (-W)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank8Mask) << (NW)
	}

	return kingAttacks
}

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
		allMoves = CreateKingAttackTable()
	case board.Knight:
		allMoves = CreateKnightAttackTable()
	default:
		panic("Invalid Piece")
	}
	for i, bb := range allMoves {
		PrintMoves(board.Square(i), bb)
	}
}
