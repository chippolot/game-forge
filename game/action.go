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
	Coord Coord
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
		Coord: Coord{
			X: x - 1,
			Y: y - 1,
		},
		Piece: getPieceFunc(),
	}, nil
}

type MovePieceAction struct {
	Start Coord
	Moves []Coord
}

func MovePieceActionDesc() ActionDesc {
	return ActionDesc{"move", "move <start_coord> <to_coord_1> ...", "moves piece at <start_coord> to <to_coord_1> and then to following coords if supported"}
}

func ParseMovePieceAction(args []string, maxMoves int) (IAction, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("invalid number of arguments for move action")
	}

	x0, y0, err := utils.ParseCoord(args[0])
	if err != nil {
		return nil, err
	}

	moves := make([]Coord, 0, 1)
	for i := 1; i < len(args); i++ {

		x, y, err := utils.ParseCoord(args[i])
		if err != nil {
			return nil, err
		}
		moves = append(moves, Coord{X: x, Y: y})
	}

	return &MovePieceAction{
		Start: Coord{
			X: x0, Y: y0,
		},
		Moves: moves,
	}, nil
}
