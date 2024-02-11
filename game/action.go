package game

import (
	"fmt"

	"github.com/chippolot/game-forge/utils"
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
	return ActionDesc{"place", "place <coord>", "places a piece at <coord>"}
}

func ParsePlacePieceAction(args []string, getPieceFunc func() Piece) (IAction, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("invalid number of arguments for place action")
	}

	x, y, err := utils.ParseCoord(args[0])
	if err != nil {
		return nil, err
	}

	return &PlacePieceAction{
		X:     x - 1,
		Y:     y - 1,
		Piece: getPieceFunc(),
	}, nil
}
