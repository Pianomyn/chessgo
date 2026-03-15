package movement

import (
	"chessgo/board"
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
	}
}

func TestRookWrapArounds(t *testing.T) {
	table := GetKingMoves()

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
