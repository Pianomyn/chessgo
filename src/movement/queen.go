package movement

import "chessgo/board"

func GetQueenAttackTable() []board.Bitboard {
	queenMoves := make([]board.Bitboard, 64)
	rookMoves := GetRookAttackTable()
	bishopMoves := GetBishopAttackTable()
	for i := range queenMoves {
		queenMoves[i] = bishopMoves[i] | rookMoves[i]
	}
	return queenMoves
}
