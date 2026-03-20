package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

func GetPawnPushes(side board.Side) []board.Bitboard {
	pushes := make([]board.Bitboard, 64)
	for i := range pushes {
		currentBit := board.Bitboard(1) << i
		if side == board.White {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.Rank2Mask != 0 {
				pushes[i] |= currentBit<<movement_utility.N | currentBit<<(movement_utility.N*2)
			} else {
				pushes[i] |= currentBit << movement_utility.N
			}
		} else {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.Rank7Mask != 0 {
				pushes[i] |= currentBit>>-movement_utility.S | currentBit>>-(movement_utility.S*2)
			} else {
				pushes[i] |= currentBit >> -movement_utility.S
			}
		}
	}
	return pushes
}

func GetPawnAttacks(side board.Side) []board.Bitboard {
	attacks := make([]board.Bitboard, 64)
	for i := range attacks {
		currentBit := board.Bitboard(1) << i
		if side == board.White {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.FileAMask != 0 {
				attacks[i] |= currentBit << movement_utility.NE
			} else if currentBit&board.FileHMask != 0 {
				attacks[i] |= currentBit << movement_utility.NW
			} else {
				attacks[i] |= currentBit<<movement_utility.NE | currentBit<<movement_utility.NW
			}
		} else {
			if currentBit&(board.Rank1Mask|board.Rank8Mask) != 0 {
				continue
			} else if currentBit&board.FileAMask != 0 {
				attacks[i] |= currentBit >> -movement_utility.SE
			} else if currentBit&board.FileHMask != 0 {
				attacks[i] |= currentBit >> -movement_utility.SW
			} else {
				attacks[i] |= currentBit>>-movement_utility.SE | currentBit>>-movement_utility.SW
			}
		}
	}
	return attacks
}
