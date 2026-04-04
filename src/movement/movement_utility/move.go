package movement_utility

import "chessgo/board"

type Move struct {
	source board.Bitboard
	target board.Square
}
