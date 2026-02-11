package main

type Colour uint8
type Piece uint8

const (
	White Colour = iota
	Black
)

const (
	Pawn Piece = iota
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

func NewChessBoard() *ChessBoard {
	cb := &ChessBoard{
		SideToMove: White,
	}

	// Pawns: Using the functional assignment style
	for i := 0; i < 8; i++ {
		offset := Square(i)
		cb.Pieces[White][Pawn] = cb.Pieces[White][Pawn].Set(8 + offset)
		cb.Pieces[Black][Pawn] = cb.Pieces[Black][Pawn].Set(48 + offset)
	}

	// Knights
	cb.Pieces[White][Knight] = cb.Pieces[White][Knight].Set(1).Set(6)
	cb.Pieces[Black][Knight] = cb.Pieces[Black][Knight].Set(57).Set(62)

	// Bishops
	cb.Pieces[White][Bishop] = cb.Pieces[White][Bishop].Set(2).Set(5)
	cb.Pieces[Black][Bishop] = cb.Pieces[Black][Bishop].Set(58).Set(61)

	// Rooks
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(0).Set(7)
	cb.Pieces[Black][Rook] = cb.Pieces[Black][Rook].Set(56).Set(63)

	// Queens
	cb.Pieces[White][Queen] = cb.Pieces[White][Queen].Set(3)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(59)

	// Kings
	cb.Pieces[White][King] = cb.Pieces[White][King].Set(4)
	cb.Pieces[Black][King] = cb.Pieces[Black][King].Set(60)

	cb.Sync()
	return cb
}

func (cb *ChessBoard) Sync() {
	cb.Colours[White] = 0
	cb.Colours[Black] = 0
	for i := range 6 {
		cb.Colours[White] |= cb.Pieces[White][i]
		cb.Colours[Black] |= cb.Pieces[Black][i]
	}

	cb.Occupied = cb.Colours[White] | cb.Colours[Black]
	cb.Empty = ^cb.Occupied
}
