package main

import (
	"fmt"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	cb := NewChessBoard()

	if cb.SideToMove != White {
		t.Error("ChessBoard initial side to move should be White")
	}

	for i := range 8 {
		offset := Square(i)
		if !cb.Pieces[White][Pawn].Get(A2 + offset) {
			t.Error("White pawns have not been placed correctly")
		}
		if !cb.Pieces[Black][Pawn].Get(A7 + offset) {
			t.Error("Black pawns have not been placed correctly")
		}
	}

	initialPositions := []struct {
		colour   Side
		piece    Piece
		position Square
	}{
		// Knights
		{White, Knight, B1},
		{White, Knight, G1},
		{Black, Knight, B8},
		{Black, Knight, G8},

		// Bishops
		{White, Bishop, C1},
		{White, Bishop, F1},
		{Black, Bishop, C8},
		{Black, Bishop, F8},

		// Rooks
		{White, Rook, A1},
		{White, Rook, H1},
		{Black, Rook, A8},
		{Black, Rook, H8},

		// Rooks
		{White, Queen, D1},
		{Black, Queen, D8},

		// Rooks
		{White, King, E1},
		{Black, King, E8},
	}

	for _, c := range initialPositions {
		t.Run(fmt.Sprintf("Position %v", c.position), func(t *testing.T) {
			if !cb.Pieces[c.colour][c.piece].Get(c.position) || !cb.Colours[c.colour].Get(c.position) {
				t.Errorf("Position %d should contain %d %d", c.position, c.colour, c.piece)
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
	if !cb.Empty.Get(A1) || cb.Empty.Get(A3) || cb.Empty.Get(H8) {
		t.Errorf("Empty bitboard incorrect")
	}
}
