package game

import (
	"fmt"
	"strconv"
)

type ActionDesc struct {
	Keyword     string
	Usage       string
	Description string
}

type IAction interface {
}

type PlacePieceAction struct {
	X, Y  int
	Piece Piece
}

func PlacePieceActionDesc() ActionDesc {
	return ActionDesc{"place", "place <x> <y>", "places a piece at <x>,<y>"}
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
