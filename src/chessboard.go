package main

type Side uint8
type Piece uint8

const (
	White Side = iota
	Black
)

func (c Side) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}

const (
	Pawn Piece = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (c Piece) String() string {
	return map[Piece]string{
		Pawn:   "Pawn",
		Knight: "Knight",
		Bishop: "Bishop",
		Rook:   "Rook",
		Queen:  "Queen",
		King:   "King",
	}[c]
}

type ChessBoard struct {
	Pieces  [2][6]Bitboard
	Colours [2]Bitboard

	Occupied   Bitboard
	Empty      Bitboard
	SideToMove Side
}

func NewChessBoard() *ChessBoard {
	cb := &ChessBoard{
		SideToMove: White,
	}

	// Pawns
	cb.Pieces[White][Pawn] |= Rank2Mask
	cb.Pieces[Black][Pawn] |= Rank7Mask

	// Knights
	cb.Pieces[White][Knight] = cb.Pieces[White][Knight].Set(B1).Set(G1)
	cb.Pieces[Black][Knight] = cb.Pieces[Black][Knight].Set(B8).Set(G8)

	// Bishops
	cb.Pieces[White][Bishop] = cb.Pieces[White][Bishop].Set(C1).Set(F1)
	cb.Pieces[Black][Bishop] = cb.Pieces[Black][Bishop].Set(C8).Set(F8)

	// Rooks
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(A1).Set(H1)
	cb.Pieces[Black][Rook] = cb.Pieces[Black][Rook].Set(A8).Set(H8)

	// Queens
	cb.Pieces[White][Queen] = cb.Pieces[White][Queen].Set(D1)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(D8)

	// Kings
	cb.Pieces[White][King] = cb.Pieces[White][King].Set(E1)
	cb.Pieces[Black][King] = cb.Pieces[Black][King].Set(E8)

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
