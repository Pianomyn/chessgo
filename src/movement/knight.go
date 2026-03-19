package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

func GetKnightMoves() []board.Bitboard {
	knightAttacks := make([]board.Bitboard, 64)
	for i := range knightAttacks {
		bit := board.Bitboard(1) << i

		knightAttacks[i] |= (bit & board.NotFileABMask) << movement_utility.NoWeWe
		knightAttacks[i] |= (bit & board.NotFileAMask) << movement_utility.NoNoWe
		knightAttacks[i] |= (bit & board.NotFileGHMask) << movement_utility.NoEaEa
		knightAttacks[i] |= (bit & board.NotFileHMask) << movement_utility.NoNoEa

		// Cant negative bitshift
		knightAttacks[i] |= (bit & board.NotFileABMask) >> (-movement_utility.SoWeWe)
		knightAttacks[i] |= (bit & board.NotFileAMask) >> (-movement_utility.SoSoWe)
		knightAttacks[i] |= (bit & board.NotFileGHMask) >> (-movement_utility.SoEaEa)
		knightAttacks[i] |= (bit & board.NotFileHMask) >> (-movement_utility.SoSoEa)
	}

	return knightAttacks
}
