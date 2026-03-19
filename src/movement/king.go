package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

func GetKingMoves() []board.Bitboard {
	kingAttacks := make([]board.Bitboard, 64)
	for i := range kingAttacks {
		bit := board.Bitboard(1) << i

		kingAttacks[i] |= (bit & board.NotRank8Mask) << (movement_utility.N)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank8Mask) << (movement_utility.NE)
		kingAttacks[i] |= (bit & board.NotFileHMask) << (movement_utility.E)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank1Mask) >> (-movement_utility.SE)
		kingAttacks[i] |= (bit & board.NotRank1Mask) >> (-movement_utility.S)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank1Mask) >> (-movement_utility.SW)
		kingAttacks[i] |= (bit & board.NotFileAMask) >> (-movement_utility.W)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank8Mask) << (movement_utility.NW)
	}

	return kingAttacks
}
