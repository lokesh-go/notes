package piece

// PlayingPiece ...
type PlayingPiece interface {
	Play() string
}

type piece struct {
	pieceType string
}

// NewPieceType ...
func NewPieceType(pieceType string) PlayingPiece {
	return &piece{
		pieceType: pieceType,
	}
}

// Play ...
func (p *piece) Play() string {
	return p.pieceType
}
