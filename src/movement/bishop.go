package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
)

type BishopRays struct {
	NE []board.Bitboard
	NW []board.Bitboard
	SE []board.Bitboard
	SW []board.Bitboard
}

/*
func GetBishopMoves(cb *board.ChessBoard, attacks BishopRays) []movement_utility.Move {
	bishops := cb.Pieces[cb.SideToMove][board.Bishop]

	for bishops != 0 {
		square := movement_utility.GetNextPiece(&bishops)

		neBlockers := attacks.NE[square] & cb.Occupied
		if neBlockers == 0 {
			// At least 1 blocker
			neHit := bits.TrailingZeros64(uint64(neBlockers))
		} else {

		}
		nwBlockers := attacks.NE[square] & cb.Occupied

	}

	return board.Bitboard(1)
}
*/

func GetBishopAttackTable() BishopRays {
	// Individual rays instead of combined so can use lsb/msb per dir
	// O(n^2), could optimize
	rays := BishopRays{
		NE: make([]board.Bitboard, 64),
		NW: make([]board.Bitboard, 64),
		SE: make([]board.Bitboard, 64),
		SW: make([]board.Bitboard, 64),
	}

	for i := range 64 {
		currentBit := board.Bitboard(1) << i
		ne := currentBit
		nw := currentBit
		se := currentBit
		sw := currentBit
		for ne != 0 || nw != 0 || se != 0 || sw != 0 {
			if ne != 0 {
				ne = (ne << movement_utility.NE) & board.NotFileAMask
				rays.NE[i] |= ne
			}
			if nw != 0 {
				nw = (nw << movement_utility.NW) & board.NotFileHMask
				rays.NW[i] |= nw
			}
			if se != 0 {
				se = (se >> -movement_utility.SE) & board.NotFileAMask
				rays.SE[i] |= se
			}
			if sw != 0 {
				sw = (sw >> -movement_utility.SW) & board.NotFileHMask
				rays.SW[i] |= sw
			}
		}
	}
	return rays
}
