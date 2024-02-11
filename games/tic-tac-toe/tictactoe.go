package tictactoe

import (
	"fmt"

	"github.com/chippolot/game-forge/game"
)

type Logic struct{}

func (l *Logic) RegisterActions(actionParser *game.ActionParser) {

	placePieceActionRegistration := game.ActionRegistration{
		Desc: game.PlacePieceActionDesc(),
		ParseFunc: func(args []string, gameInstance game.IGame) (game.IAction, error) {
			return game.ParsePlacePieceAction(args, func() game.Piece {
				return getPlayerPiece(gameInstance.GetState().GetCurrentPlayer())
			})
		},
	}
	actionParser.RegisterAction(placePieceActionRegistration)
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
	// Check rows and columns
	for i := 0; i < 3; i++ {
		hasWinner, winner := checkRun(board, 0, i, 1, 0)
		if hasWinner {
			return true, winner
		}
		hasWinner, winner = checkRun(board, i, 0, 0, 1)
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
	game.ICommonGameState
}

func NewState(board game.IBoard) game.IGameState {
	return &GameState{game.NewCommonGameState(board)}
}

type Piece struct {
	player game.Player
}

func NewGame(parser *game.ActionParser) game.IGame {
	name := "Tic-Tac-Toe"
	desc := "Tic-tac-toe is a classic two-player game played on a 3x3 grid. Players take turns marking spaces with their " +
		"respective symbols, typically \"X\" and \"O\", with the objective of placing three of their symbols in a row, column," +
		" or diagonal. The first player to achieve this goal wins the game. If all spaces are filled without a winner, the game " +
		"ends in a draw. Tic-tac-toe is easy to learn, yet offers strategic depth, making it a timeless and engaging pastime for players of all ages."
	metadata := game.NewMetadata(name, desc)
	logic := &Logic{}
	board := game.NewBoard(3, 3)
	state := NewState(board)
	renderer := &game.SimpleGameRenderer{PrintScores: false}
	logic.RegisterActions(parser)
	return game.NewGame(logic, state, renderer, metadata, parser)
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

func getPlayerPiece(player game.Player) game.Piece {
	return Piece{
		player: player,
	}
}
