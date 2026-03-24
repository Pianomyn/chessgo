package movement

import "chessgo/board"

func GetQueenAttacks() []board.Bitboard {
	queenMoves := make([]board.Bitboard, 64)
	rookMoves := GetRookAttacks()
	bishopMoves := GetBishopAttacks()
	for i := range queenMoves {
		queenMoves[i] = bishopMoves[i] | rookMoves[i]
	}
	return queenMoves
}
