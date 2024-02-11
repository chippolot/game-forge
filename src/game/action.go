package game

import (
	"fmt"
	"strconv"
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

func ParsePlacePieceAction(args []string, getPieceFunc func() Piece) (IAction, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid number of arguments for place action")
	}

	x, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse X coordinate: %w", err)
	}

	y, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Y coordinate: %w", err)
	}

	return &PlacePieceAction{
		X:     x - 1,
		Y:     y - 1,
		Piece: getPieceFunc(),
	}, nil
}
