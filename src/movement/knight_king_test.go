package movement

import (
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
