package movement_utility

import (
	"chessgo/board"
	"math/bits"
)

func RayMoves(ray board.Bitboard, allRays []board.Bitboard, occupied board.Bitboard, positive bool) board.Bitboard {
	blockers := ray & occupied
	if blockers == 0 {
		return ray
	}

	var first int
	if positive {
		first = bits.TrailingZeros64(uint64(blockers))
	} else {
		first = 63 - bits.LeadingZeros64(uint64(blockers))
	}

	return ray ^ allRays[first]
}

type Move struct {
	Source      board.Square
	Target      board.Square
	SourcePiece board.Piece

	// Promotion, castling, en passant etc can be handled later
}
