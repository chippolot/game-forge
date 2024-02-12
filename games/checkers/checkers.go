package checkers

import (
	"fmt"

	"github.com/chippolot/game-forge/game"
)

type Logic struct{}

func (l *Logic) RegisterActions(actionParser *game.ActionParser) {
	movePieceActionRegistration := game.ActionRegistration{
		Desc: game.ActionDesc{
			Keyword:     "move",
			Usage:       "move <fromX> <fromY> to <toX> <toY>",
			Description: "Move a piece from coordinates <fromX>,<fromY> to <toX>,<toY>, capturing opponent's pieces if applicable.",
		},
		ParseFunc: func(args []string, gameInstance game.IGame) (game.IAction, error) {
			return game.ParseMovePieceAction(args, -1)
		},
	}
	actionParser.RegisterAction(movePieceActionRegistration)
}

func (l *Logic) ExecuteAction(action game.IAction, state game.IGameState) (game.GameResult, error) {
	validAction, err := isValidAction(action, state)
	if !validAction {
		return game.GameResult{}, err
	}

	switch typedAction := action.(type) {
	case *game.MovePieceAction:
		executeMove(typedAction.Start, typedAction.Moves[0], state)
	default:
		panic("Invalid action.")
	}

	updatePieces(state.GetBoard())
	gameOverState, winningPlayer := isGameOver(state)
	if gameOverState == game.NotGameOver {
		currentPlayer := state.GetCurrentPlayer()
		currentPlayer = (currentPlayer + 1) % 2
		state.SetCurrentPlayer(currentPlayer)
	}

	gameResult := game.GameResult{
		State:         gameOverState,
		WinningPlayer: winningPlayer,
	}
	return gameResult, nil
}

func isValidAction(action game.IAction, state game.IGameState) (bool, error) {
	switch typedAction := action.(type) {
	case *game.MovePieceAction:
		return isValidMove(typedAction.Start, typedAction.Moves[0], state), nil
	default:
		return false, fmt.Errorf("unsupported action %T", typedAction)
	}
}

func isGameOver(state game.IGameState) (game.GameResultState, game.Player) {
	if noMoreMoves(state) {
		return game.GameWon, calculateWinningPlayer(state)
	}
	return game.NotGameOver, 0
}

func noMoreMoves(state game.IGameState) bool {
	// Implement logic to check if there are no more legal moves for the current player
	return false // Placeholder return value
}

func calculateWinningPlayer(state game.IGameState) game.Player {
	// Implement logic to determine the winning player based on remaining pieces
	return 0 // Placeholder return value
}

func isValidMove(from game.Coord, to game.Coord, state game.IGameState) bool {
	// Implement logic to validate move based on Checkers rules
	return true // Placeholder return value
}

func executeMove(from game.Coord, to game.Coord, state game.IGameState) {
	// Implement logic to execute the move, including capturing and "kinging" pieces as applicable
}

func updatePieces(board game.IBoard) {
	// Implement logic to update pieces after each move, potentially removing captured pieces and updating piece states (e.g., "king" status)
}

type GameState struct {
	game.ICommonGameState
}

func NewState(board game.IBoard) game.IGameState {
	s := &GameState{game.NewCommonGameState(board)}
	initializeBoard(s.GetBoard())
	return s
}

func initializeBoard(board game.IBoard) {
	// Set up the Checkers board with pieces in starting positions
}

type Piece struct {
	player game.Player
	king   bool // Indicates whether the piece has been "kinged"
}

func (p Piece) GetPlayer() game.Player {
	return p.player
}

func (p Piece) IsKing() bool {
	return p.king
}

func NewGame(parser *game.ActionParser) game.IGame {
	name := "Checkers"
	desc := "Checkers is a strategy board game for two players which involves diagonal moves of uniform game pieces and mandatory captures by jumping over opponent pieces."
	metadata := game.NewMetadata(name, desc)
	logic := &Logic{}
	board := game.NewBoard(8, 8)
	state := NewState(board)
	renderer := &game.SimpleGameRenderer{PrintScores: true}
	logic.RegisterActions(parser)
	return game.NewGame(logic, state, renderer, metadata, parser)
}
