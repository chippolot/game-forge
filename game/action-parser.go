package game

import (
	"fmt"
	"strings"
)

type ActionKeyword string

type ActionRegistration struct {
	Desc      ActionDesc
	ParseFunc ActionParseFunc
}

type ActionParseFunc func([]string, IGame) (IAction, error)

type ActionParser struct {
	registry map[string]ActionRegistration
}

func NewActionParser() *ActionParser {
	return &ActionParser{
		registry: make(map[string]ActionRegistration),
	}
}

func (p *ActionParser) PrintAvailableActions() {
	fmt.Println("Available Actions:")
	for _, reg := range p.registry {
		fmt.Printf("%v\tusage: %v\n", reg.Desc.Keyword, reg.Desc.Usage)
	}
}

// RegisterAction registers an action type.
func (p *ActionParser) RegisterAction(registration ActionRegistration) {
	p.registry[registration.Desc.Keyword] = registration
}

// ParseAction parses an action from the input string.
func (p *ActionParser) ParseAction(input string, gameInstance IGame) (IAction, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	actionType := parts[0]
	registration, found := p.registry[actionType]
	if !found {
		return nil, fmt.Errorf("invalid action: %s", actionType)
	}

	return registration.ParseFunc(parts[1:], gameInstance)
}
