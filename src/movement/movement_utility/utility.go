package movement_utility

import (
	"chessgo/board"
	"math/bits"
)

func GetNextPiece(pieces *board.Bitboard) int {
	/*
		For a bitboard containing locations of all pieces of a particular type,
		yield the location of the next piece
	*/
	if *pieces == 0 {
		return -1
	}

	lsbIndex := bits.TrailingZeros64(uint64(*pieces))

	pieces.Clear(board.Square(lsbIndex))
	return lsbIndex
}
