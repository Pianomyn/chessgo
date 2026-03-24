package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

func GetBishopAttacks() []board.Bitboard {
	// O(n^2), could optimize
	bishopMoves := make([]board.Bitboard, 64)
	for i := range bishopMoves {
		currentBit := board.Bitboard(1) << i
		ne := currentBit
		nw := currentBit
		se := currentBit
		sw := currentBit
		for ne != 0 || nw != 0 || se != 0 || sw != 0 {
			if ne != 0 {
				ne = (ne << movement_utility.NE) & board.NotFileAMask
				bishopMoves[i] |= ne
			}
			if nw != 0 {
				nw = (nw << movement_utility.NW) & board.NotFileHMask
				bishopMoves[i] |= nw
			}
			if se != 0 {
				se = (se >> -movement_utility.SE) & board.NotFileAMask
				bishopMoves[i] |= se
			}
			if sw != 0 {
				sw = (sw >> -movement_utility.SW) & board.NotFileHMask
				bishopMoves[i] |= sw
			}
		}
		bishopMoves[i] &^= currentBit
	}
	return bishopMoves
}
