package game

type Piece interface {
	GetPlayer() Player
	GetDisplayString() string
}
