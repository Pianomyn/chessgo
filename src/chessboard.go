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

func NewChessBoard() *ChessBoard {
	cb := &ChessBoard{
		SideToMove: White,
	}

	// Pawns
	for i := 0; i < 8; i++ {
		cb.Pieces[White][Pawn].Set(8 + i)  // Rank 2
		cb.Pieces[Black][Pawn].Set(48 + i) // Rank 7
	}

	// Knights
	cb.Pieces[White][Knight].Set(1)
	cb.Pieces[White][Knight].Set(6)
	cb.Pieces[Black][Knight].Set(57)
	cb.Pieces[Black][Knight].Set(62)

	// Bishops
	cb.Pieces[White][Bishop].Set(2)
	cb.Pieces[White][Bishop].Set(5)
	cb.Pieces[Black][Bishop].Set(58)
	cb.Pieces[Black][Bishop].Set(61)

	// Rooks
	cb.Pieces[White][Rook].Set(0)
	cb.Pieces[White][Rook].Set(7)
	cb.Pieces[Black][Rook].Set(56)
	cb.Pieces[Black][Rook].Set(63)

	// Queens
	cb.Pieces[White][Queen].Set(3)
	cb.Pieces[Black][Queen].Set(59)

	// Kings
	cb.Pieces[White][King].Set(4)
	cb.Pieces[Black][King].Set(60)

	cb.Sync()
	return cb
}

func (cb *ChessBoard) Sync() {
	for i := range 6 {
		cb.Colours[White] |= cb.Pieces[White][i]
		cb.Colours[Black] |= cb.Pieces[Black][i]
	}

	cb.Occupied = cb.Colours[White] | cb.Colours[Black]
	cb.Empty = ^cb.Occupied
}
