package movement

import (
	"chessgo/board"
	"math/bits"
	"testing"
)

func TestKnightAttackCounts(t *testing.T) {
	table := CreateKnightAttackTable()

	tests := []struct {
		name     string
		square   int
		expected int
	}{
		{"Corner (a1)", 0, 2},
		{"Edge (a2)", 8, 3},
		{"Near Edge (b2)", 9, 4},
		{"Outer Ring (c2)", 10, 6},
		{"Center (d4)", 27, 8},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestKnightWrapArounds(t *testing.T) {
	table := CreateKnightAttackTable()

	tests := []struct {
		name           string
		square         int
		forbiddenMasks board.Bitboard
	}{
		{"a1", 0, board.FileGMask | board.FileHMask},
		{"b1", 1, board.FileHMask},
		{"h8", 63, board.FileAMask | board.FileBMask},
		{"g8", 62, board.FileAMask},
	}

	for _, tt := range tests {
		illegalMoves := table[tt.square] & tt.forbiddenMasks
		if illegalMoves != 0 {
			t.Errorf("%s (square %d): detected wrap-around.", tt.name, tt.square)
		}
	}
}

func TestKingAttackCounts(t *testing.T) {
	table := CreateKingAttackTable()
	tests := []struct {
		name     string
		square   int
		expected int
	}{

		{"Corner (a1)", 0, 3},
		{"Edge (a2)", 8, 5},
		{"Center (d4)", 27, 8},
	}

	for _, tt := range tests {
		got := bits.OnesCount64(uint64(table[tt.square]))
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func TestKingWrapArounds(t *testing.T) {
	table := CreateKingAttackTable()

	tests := []struct {
		name           string
		square         int
		forbiddenMasks board.Bitboard
	}{
		{"a1", 0, board.FileHMask},
		{"a4", 24, board.FileHMask},
		{"h1", 7, board.FileAMask},
		{"h8", 63, board.FileAMask},
	}

	for _, tt := range tests {
		illegalMoves := table[tt.square] & tt.forbiddenMasks
		if illegalMoves != 0 {
			t.Errorf("%s (square %d): detected wrap-around.", tt.name, tt.square)
		}
	}
}
