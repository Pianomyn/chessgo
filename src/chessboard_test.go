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
		if !cb.Pieces[White][Pawn].Get(8 + offset) {
			t.Error("White pawns have not been placed correctly")
		}
		if !cb.Pieces[Black][Pawn].Get(48 + offset) {
			t.Error("Black pawns have not been placed correctly")
		}
	}

	initialPositions := []struct {
		colour   Colour
		piece    Piece
		position Square
	}{
		// Knights
		{White, Knight, 1},
		{White, Knight, 6},
		{Black, Knight, 57},
		{Black, Knight, 62},

		// Bishops
		{White, Bishop, 2},
		{White, Bishop, 5},
		{Black, Bishop, 58},
		{Black, Bishop, 61},

		// Rooks
		{White, Rook, 0},
		{White, Rook, 7},
		{Black, Rook, 56},
		{Black, Rook, 63},

		// Rooks
		{White, Queen, 3},
		{Black, Queen, 59},

		// Rooks
		{White, King, 4},
		{Black, King, 60},
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
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Clear(0)
	cb.Pieces[White][Rook] = cb.Pieces[White][Rook].Set(16)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Clear(63)
	cb.Pieces[Black][Queen] = cb.Pieces[Black][Queen].Set(63)

	cb.Sync()

	if cb.Colours[White].Get(0) {
		t.Errorf("White colour bitboard didn't clear rook")
	}
	if !cb.Colours[White].Get(16) {
		t.Errorf("White colour bitboard didn't place rook")
	}
	if !cb.Colours[Black].Get(63) {
		t.Errorf("Black colour bitboard missing queen square")
	}
	if cb.Occupied.Get(0) || !cb.Occupied.Get(16) || !cb.Occupied.Get(63) {
		t.Errorf("Occupied bitboard incorrect")
	}
	if !cb.Empty.Get(0) || cb.Empty.Get(16) || cb.Empty.Get(63) {
		t.Errorf("Empty bitboard incorrect")
	}
}
