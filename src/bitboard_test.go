package main

import (
	"testing"
)

func TestBitboardManipulation(t *testing.T) {
	cases := []struct {
		name   string
		square int
	}{
		{"a1", 0},
		{"h1", 7},
		{"a8", 56},
		{"h8", 63},
		{"middle", 32},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var bb Bitboard

			bb.Set(c.square)
			if !bb.Get(c.square) {
				t.Errorf("%s bit not set\n", c.name)
			}

			bb.Clear(c.square)
			if bb.Get(c.square) {
				t.Errorf("%s bit not cleared\n", c.name)
			}
		})
	}
}
