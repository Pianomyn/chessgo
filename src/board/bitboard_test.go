package board

import (
	"fmt"
	"testing"
)

func TestBitboardManipulation(t *testing.T) {
	cases := []struct {
		square Square
	}{
		{A1},
		{H1},
		{A8},
		{H8},
		{A5},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s", c.square), func(t *testing.T) {
			var bb Bitboard

			bb = bb.Set(c.square)
			if !bb.Get(c.square) {
				t.Errorf("%s bit not set\n", c.square)
			}

			bb = bb.Clear(c.square)
			if bb.Get(c.square) {
				t.Errorf("%s bit not cleared\n", c.square)
			}
		})
	}
}
