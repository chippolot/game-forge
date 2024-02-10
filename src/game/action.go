package game

import (
	"fmt"
)

const (
	PlacePieceActionKeyword = "place"
)

type IAction interface {
	GetKeyword() ActionKeyword
	Describe() string
}

type PlacePieceAction struct {
	X, Y  int
	Piece Piece
}

func (a PlacePieceAction) GetKeyword() ActionKeyword {
	return PlacePieceActionKeyword
}

func (a PlacePieceAction) Describe() string {
	return fmt.Sprintf("%v x y", a.GetKeyword())
}
