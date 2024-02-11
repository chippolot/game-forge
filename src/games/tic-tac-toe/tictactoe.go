package tictactoe

import (
	"fmt"
	"strconv"

	"github.com/chippolot/game-forge/src/game"
)

type Logic struct{}

func (l *Logic) RegisterActions(actionParser *game.ActionParser) {
	actionParser.RegisterAction(game.PlacePieceActionKeyword, parsePlacePieceAction)
}

func (l *Logic) ExecuteAction(action game.IAction, state game.IGameState) (game.GameResult, error) {
	validAction, err := isValidAction(action, state)
	if !validAction {
		return game.GameResult{}, err
	}

	switch typedAction := action.(type) {
	case *game.PlacePieceAction:
		state.GetBoard().PlacePiece(typedAction.X, typedAction.Y, typedAction.Piece)
	default:
		panic("Invalid action.")
	}

	gameOverState, winningPlayer := isGameOver(state.GetBoard())
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
	case *game.PlacePieceAction:
		x, y := typedAction.X, typedAction.Y
		if !state.GetBoard().IsInBounds(x, y) {
			return false, fmt.Errorf("out of bounds")
		}
		if state.GetBoard().GetPiece(x, y) != nil {
			return false, fmt.Errorf("space occupied")
		}
		return true, nil
	default:
		return false, fmt.Errorf("unsupported action %T", typedAction)
	}
}

func isGameOver(board game.IBoard) (game.GameResultState, game.Player) {
	hasWinner, winningPlayer := getWinner(board)
	if hasWinner {
		return game.GameWon, winningPlayer
	}
	if board.IsFull() {
		return game.GameTie, 0
	}
	return game.NotGameOver, 0
}

func getWinner(board game.IBoard) (bool, game.Player) {
	// Check columns
	for x := 0; x < 3; x++ {
		hasWinner, winner := checkRun(board, x, 0, 0, 1)
		if hasWinner {
			return true, winner
		}
	}
	// Check rows
	for y := 0; y < 3; y++ {
		hasWinner, winner := checkRun(board, 0, y, 1, 0)
		if hasWinner {
			return true, winner
		}
	}
	// Check diagonals
	hasWinner, winner := checkRun(board, 0, 0, 1, 1)
	if hasWinner {
		return true, winner
	}
	hasWinner, winner = checkRun(board, 2, 0, -1, 1)
	if hasWinner {
		return true, winner
	}
	return false, 0
}

func checkRun(board game.IBoard, x, y, dx, dy int) (bool, game.Player) {
	piece := board.GetPiece(x, y)
	if piece == nil {
		return false, 0
	}
	for i := 0; i < 2; i++ {
		x += dx
		y += dy
		nextPiece := board.GetPiece(x, y)
		if nextPiece == nil || nextPiece.GetPlayer() != piece.GetPlayer() {
			return false, 0
		}
	}
	return true, piece.GetPlayer()
}

type GameState struct {
	game.CommonGameState
}

func NewState() game.IGameState {
	return &GameState{}
}

// Game concrete implementation of the tic-tac-toe game
type Game struct {
	game.Metadata
	logic game.ILogic
	state game.IGameState
}

func NewGame(parser *game.ActionParser) game.IGame {
	name := "Tic-Tac-Toe"
	desc := "Tic-tac-toe is a classic two-player game played on a 3x3 grid. Players take turns marking spaces with their respective symbols, typically \"X\" and \"O\", with the objective of placing three of their symbols in a row, column, or diagonal. The first player to achieve this goal wins the game. If all spaces are filled without a winner, the game ends in a draw. Tic-tac-toe is easy to learn, yet offers strategic depth, making it a timeless and engaging pastime for players of all ages."
	game := &Game{
		Metadata: *game.NewMetadata(name, desc),
		logic:    &Logic{},
		state:    NewState(),
	}
	game.logic.RegisterActions(parser)
	return game
}

func (g *Game) Start() {
	g.currentPlayer = 0
}

func (g *Game) GetCurrentPlayer() game.Player {
	return g.currentPlayer
}

func (g *Game) GetPlayerPiece(player game.Player) game.Piece {
	return Piece{
		player: player,
	}
}

func (g *Game) ExecuteAction(action game.IAction) (game.GameResult, error) {
	return g.logic.ExecuteAction(action, g.state)
}

func (g *Game) Restart() {
	g.gameBoard.Clear()
	g.Start()
}

type Piece struct {
	player game.Player
}

func (p Piece) GetPlayer() game.Player {
	return p.player
}

func (p Piece) GetDisplayString() string {
	if p.player == 0 {
		return "x"
	} else if p.player == 1 {
		return "o"
	}
	panic("unknown piece")
}

func parsePlacePieceAction(args []string, gameInstance game.IGame) (game.IAction, error) {
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

	return &game.PlacePieceAction{
		X:     x - 1,
		Y:     y - 1,
		Piece: gameInstance.GetPlayerPiece(gameInstance.GetCurrentPlayer()),
	}, nil
}
