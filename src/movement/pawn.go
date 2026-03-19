package movement

import "chessgo/board"

func GetPawnMoves() []board.Bitboard {
	pushes := GetPawnPushes()
	attacks := GetPawnAttacks()
	moves := make([]board.Bitboard, 64)
	for i := range pushes {
		moves[i] = pushes[i] | attacks[i]
	}
	return moves
}

func GetPawnPushes() []board.Bitboard {
	pushes := make([]board.Bitboard, 64)
	return pushes

}

func GetPawnAttacks() []board.Bitboard {
	attacks := make([]board.Bitboard, 64)
	return attacks
}
