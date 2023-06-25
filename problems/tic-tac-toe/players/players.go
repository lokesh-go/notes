package players

import "github.com/lokesh-go/notes/problems/tic-tac-toe/piece"

// Player ...
type Player struct {
	Name         string
	PlayingPiece piece.PlayingPiece
}

// New ...
func New(name, pieceType string) Player {
	return Player{
		Name:         name,
		PlayingPiece: piece.NewPieceType(pieceType),
	}
}
