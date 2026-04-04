package board

import (
	"fmt"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	cb := NewChessBoard()

	if cb.SideToMove != White {
		t.Error("ChessBoard initial side to move should be White")
	}

	pieceTests := []struct {
		pos      Square
		expected ColouredPiece
	}{
		{A1, WhiteRook}, {B1, WhiteKnight}, {C1, WhiteBishop}, {D1, WhiteQueen},
		{E1, WhiteKing}, {F1, WhiteBishop}, {G1, WhiteKnight}, {H1, WhiteRook},
		{A8, BlackRook}, {B8, BlackKnight}, {C8, BlackBishop}, {D8, BlackQueen},
		{E8, BlackKing}, {F8, BlackBishop}, {G8, BlackKnight}, {H8, BlackRook},
	}
	for i := 0; i < 8; i++ {
		pieceTests = append(pieceTests,
			struct {
				pos      Square
				expected ColouredPiece
			}{A2 + Square(i), WhitePawn},
			struct {
				pos      Square
				expected ColouredPiece
			}{A7 + Square(i), BlackPawn},
		)
	}

	for _, tt := range pieceTests {
		t.Run(fmt.Sprintf("Square_%v", tt.pos), func(t *testing.T) {
			if cb.Mailbox[tt.pos] != tt.expected {
				t.Errorf("Mailbox at %v: expected %v, got %v", tt.pos, tt.expected, cb.Mailbox[tt.pos])
			}

			if tt.expected != NoPiece {
				side := Side(uint8(tt.expected) / 6)
				piece := Piece(uint8(tt.expected) % 6)

				if !cb.Pieces[side][piece].Get(tt.pos) {
					t.Errorf("Bitboard for %v missing at %v", piece, tt.pos)
				}
				if !cb.Colours[side].Get(tt.pos) {
					t.Errorf("Colour bitboard for %v missing at %v", side, tt.pos)
				}
			}
		})
	}
}

func TestSyncBasic(t *testing.T) {
	cb := &ChessBoard{}
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Clear(A1)
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(A3)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Clear(H8)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(H8)

	cb.Sync()

	if cb.Colours[White].Get(A1) {
		t.Errorf("White colour bitboard didn't clear rook")
	}
	if !cb.Colours[White].Get(A3) {
		t.Errorf("White colour bitboard didn't place rook")
	}
	if !cb.Colours[Black].Get(H8) {
		t.Errorf("Black colour bitboard missing queen square")
	}
	if cb.Occupied.Get(A1) || !cb.Occupied.Get(A3) || !cb.Occupied.Get(H8) {
		t.Errorf("Occupied bitboard incorrect")
	}
}
