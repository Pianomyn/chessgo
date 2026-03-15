package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestRookMovement(t *testing.T) {
	table := GetRookMoves()
	tests := []struct {
		name     string
		square   board.Square
		expected board.Bitboard
	}{

		{"a1", board.A1, (board.FileAMask | board.Rank1Mask).Clear(board.A1)},
		{"a2", board.A2, (board.FileAMask | board.Rank2Mask).Clear(board.A2)},
		{"b2", board.B2, (board.FileBMask | board.Rank2Mask).Clear(board.B2)},
		{"d4", board.D4, (board.FileDMask | board.Rank4Mask).Clear(board.D4)},
	}

	for _, tt := range tests {
		if table[tt.square] != tt.expected {
			t.Errorf("%s: Mismatch between expected and actual bitboard", tt.name)
		}
		if bits.OnesCount64(uint64(table[tt.square])) != 14 {
			t.Errorf("%s: Number of set bits was not 14", tt.name)
		}
	}
}
