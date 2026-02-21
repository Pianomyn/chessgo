package movement

import (
	"chessgo/board"
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
		got := countSetBits(table[tt.square])
		if got != tt.expected {
			t.Errorf("%s: expected %d attacks, got %d", tt.name, tt.expected, got)
		}
	}
}

func countSetBits(b board.Bitboard) int {
	count := 0
	for b != 0 {
		b &= (b - 1)
		count++
	}
	return count
}
