package main

import "fmt"

type Bitboard uint64

const (
	EmptyBitboard Bitboard = 0
	FullBitboard           = ^EmptyBitboard
)

func (b Bitboard) String() string {
	var s string
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if (b>>square)&1 == 1 {
				s += "X "
			} else {
				s += ". "
			}
		}
		s += "\n"
	}
	return s
}

func printBitboards() {
	fmt.Println(FullBitboard)
	fmt.Println(EmptyBitboard)
}
