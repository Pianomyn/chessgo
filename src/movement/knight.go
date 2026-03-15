package movement

import "chessgo/board"

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

func GetKnightMoves() []board.Bitboard {
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
