package main

import (
	"fmt"
	"strings"
)

type Bitboard uint64
type Square uint8

const (
	EmptyBitboard Bitboard = 0
	FullBitboard           = ^EmptyBitboard
)

func (b *Bitboard) String() string {
	var sb strings.Builder
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if (*b>>square)&1 == 1 {
				sb.WriteString("X ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b *Bitboard) Set(square Square) {
	*b |= (1 << square)
}

func (b *Bitboard) Clear(square Square) {
	*b &= ^(1 << square)
}

func (b *Bitboard) Get(square Square) bool {
	return (*b>>square)&1 == 1
}

func printBitboards() {
	fmt.Println(FullBitboard)
	fmt.Println(EmptyBitboard)
}
