package movement_utility

import "chessgo/board"

type Move struct {
	Source      board.Square
	Target      board.Square
	SourcePiece board.Piece
	TargetPiece board.Piece

	// Promotion, castling, en passant etc can be handled later
}
