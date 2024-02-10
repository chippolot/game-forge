package game

import (
	"fmt"
	"strings"
)

type ActionKeyword string

const (
	PlacePieceActionKeyword = "place"
)

type IAction interface {
	GetKeyword() ActionKeyword
	Describe() string
}

type ActionParser struct {
	registeredActions map[string]ActionParseFunc
}

type ActionParseFunc func([]string, IGame) (IAction, error)

func NewActionParser() *ActionParser {
	return &ActionParser{
		registeredActions: make(map[string]ActionParseFunc),
	}
}

// RegisterAction registers an action type.
func (parser *ActionParser) RegisterAction(actionType string, actionFunc ActionParseFunc) {
	parser.registeredActions[actionType] = actionFunc
}

// ParseAction parses an action from the input string.
func (parser *ActionParser) ParseAction(input string, gameInstance IGame) (IAction, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	actionType := parts[0]
	actionFunc, found := parser.registeredActions[actionType]
	if !found {
		return nil, fmt.Errorf("unknown action: %s", actionType)
	}

	return actionFunc(parts[1:], gameInstance)
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
