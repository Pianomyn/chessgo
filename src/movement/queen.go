package movement

import "chessgo/board"

func GetQueenMoves() []board.Bitboard {
	queenMoves := make([]board.Bitboard, 64)
	rookMoves := GetRookMoves()
	bishopMoves := GetBishopMoves()
	for i := range queenMoves {
		queenMoves[i] = bishopMoves[i] | rookMoves[i]
	}
	return queenMoves
}
