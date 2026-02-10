package main

type Colour uint8

const (
	White Colour = iota
	Black
)

const (
	Pawn uint8 = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type ChessBoard struct {
	Pieces  [2][6]Bitboard
	Colours [2]Bitboard

	Occupied   Bitboard
	Empty      Bitboard
	SideToMove Colour
}
