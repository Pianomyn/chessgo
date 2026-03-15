package movement

import (
	"chessgo/board"
)

func GetKingMoves() []board.Bitboard {
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
