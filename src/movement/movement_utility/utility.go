package movement_utility

import (
	"chessgo/board"
	"math/bits"
)

func GetNextBitIndex(b *board.Bitboard) int {
	if *b == 0 {
		return -1
	}

	lsbIndex := bits.TrailingZeros64(uint64(*b))

	b.Clear(board.Square(lsbIndex))
	return lsbIndex
}
