package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestKnightAttackCounts(t *testing.T) {
	table := GetKnightMoves()

	tests := []struct {
		name     string
		square   board.Square
		expected int
	}{
		{"Corner (a1)", board.A1, 2},
		{"Edge (a2)", board.A2, 3},
		{"Near Edge (b2)", board.B2, 4},
		{"Outer Ring (c2)", board.C2, 6},
		{"Center (d4)", board.D4, 8},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestKnightWrapArounds(t *testing.T) {
	table := GetKnightMoves()

	tests := []struct {
		name           string
		square         board.Square
		forbiddenMasks board.Bitboard
	}{
		{"a1", board.A1, board.FileGMask | board.FileHMask},
		{"b1", board.B1, board.FileHMask},
		{"h8", board.H8, board.FileAMask | board.FileBMask},
		{"g8", board.G8, board.FileAMask},
	}

	for _, tt := range tests {
		illegalMoves := table[tt.square] & tt.forbiddenMasks
		if illegalMoves != 0 {
			t.Errorf("%s (square %d): detected wrap-around.", tt.name, tt.square)
		}
	}
}
