package board

import (
	"fmt"
	"strings"
)

type Bitboard uint64
type Square uint8

func (s Square) File() uint8 { return uint8(s) & 7 }
func (s Square) Rank() uint8 { return uint8(s) >> 3 }

func (s Square) String() string {
	if s >= 64 {
		return "-"
	}
	return fmt.Sprintf("%c%d", 'a'+s.File(), s.Rank()+1)
}

func (b Bitboard) Set(square Square) Bitboard {
	return b | (1 << square)
}

func (b Bitboard) Clear(square Square) Bitboard {
	return b & ^(1 << square)
}

func (b Bitboard) Get(square Square) bool {
	return (b>>square)&1 == 1
}

func (b Bitboard) String() string {
	var sb strings.Builder
	for rank := 7; rank >= 0; rank-- {
		for file := range 8 {
			if b.Get(Square(rank*8 + file)) {
				sb.WriteString("X ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func printBitboards() {
	fmt.Println(FullBitboard)
	fmt.Println(EmptyBitboard)
}

const (
	EmptyBitboard Bitboard = 0
	FullBitboard           = ^EmptyBitboard
)

const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)
