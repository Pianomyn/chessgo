package movement

import (
	"chessgo/board"
	"chessgo/movement/movement_utility"
	"math/bits"
)

func GetKingMoves(cb *board.ChessBoard) []movement_utility.Move {
	attackTable := GetKingAttackTable()
	var moves []movement_utility.Move

	source := board.Square(bits.TrailingZeros64(uint64(cb.Pieces[cb.SideToMove][board.King])))
	targets := attackTable[source] &^ cb.Colours[cb.SideToMove]
	colouredKing := board.PieceToColouredPiece(board.King, cb.SideToMove)

	for targets != 0 {
		target := board.GetNextPieceSquare(&targets)
		targetPiece := cb.Mailbox[target]
		moves = append(moves, movement_utility.Move{
			Source:        source,
			Target:        target,
			SourcePiece:   colouredKing,
			CapturedPiece: targetPiece,
			Flags:         movement_utility.MoveFlags{},
		})
	}

	// Castling
	if cb.SideToMove == board.White {
		if cb.CastlingRights.WhiteKingside {

		}
		if cb.CastlingRights.WhiteQueenside {

		}
	} else {
		if cb.CastlingRights.BlackKingside {

		}
		if cb.CastlingRights.BlackQueenside {

		}
	}

	return moves
}

func GetKingAttackTable() []board.Bitboard {
	kingAttacks := make([]board.Bitboard, 64)
	for i := range kingAttacks {
		bit := board.Bitboard(1) << i

		kingAttacks[i] |= (bit & board.NotRank8Mask) << (movement_utility.N)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank8Mask) << (movement_utility.NE)
		kingAttacks[i] |= (bit & board.NotFileHMask) << (movement_utility.E)
		kingAttacks[i] |= (bit & board.NotFileHMask & board.NotRank1Mask) >> (-movement_utility.SE)
		kingAttacks[i] |= (bit & board.NotRank1Mask) >> (-movement_utility.S)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank1Mask) >> (-movement_utility.SW)
		kingAttacks[i] |= (bit & board.NotFileAMask) >> (-movement_utility.W)
		kingAttacks[i] |= (bit & board.NotFileAMask & board.NotRank8Mask) << (movement_utility.NW)
	}

	return kingAttacks
}
