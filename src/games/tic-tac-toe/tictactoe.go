package tictactoe

import (
	"fmt"
	"strconv"

	"github.com/chippolot/game-forge/src/game"
)

// Rules concrete implementation of the tic-tac-toe game rules
type Rules struct{}

func (r *Rules) IsValidAction(action game.IAction, player game.Player, board game.IBoard) (bool, error) {
	switch typedAction := action.(type) {
	case *game.PlacePieceAction:
		x, y := typedAction.X, typedAction.Y
		if !board.IsInBounds(x, y) {
			return false, fmt.Errorf("out of bounds")
		}
		if board.GetPiece(x, y) != nil {
			return false, fmt.Errorf("space occupied")
		}
		return true, nil
	default:
		return false, fmt.Errorf("unsupported action %T", typedAction)
	}
}

func (r *Rules) IsGameOver(board game.IBoard) (game.GameOverState, game.Player) {
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

// Game concrete implementation of the tic-tac-toe game
type Game struct {
	gameBoard     game.IBoard
	rules         game.IRules
	currentPlayer game.Player
}

func NewGame() game.IGame {
	return &Game{
		gameBoard:     game.NewBoard(3, 3),
		rules:         &Rules{},
		currentPlayer: 0,
	}
}

func (g *Game) GetName() string {
	return "Tic-Tac-Toe"
}

func (g *Game) GetDescription() string {
	return "Tic-tac-toe is a classic two-player game played on a 3x3 grid. Players take turns marking spaces with their respective symbols, typically \"X\" and \"O\", with the objective of placing three of their symbols in a row, column, or diagonal. The first player to achieve this goal wins the game. If all spaces are filled without a winner, the game ends in a draw. Tic-tac-toe is easy to learn, yet offers strategic depth, making it a timeless and engaging pastime for players of all ages."
}

func (g *Game) GetRules() game.IRules {
	return g.rules
}

func (g *Game) GetBoard() game.IBoard {
	return g.gameBoard
}

func (g *Game) Start() {
	g.currentPlayer = 0
}

func (g *Game) RegisterActions(actionParser *game.ActionParser) {
	actionParser.RegisterAction(game.PlacePieceActionKeyword, parsePlacePieceAction)
}

func (g *Game) GetCurrentPlayer() game.Player {
	return g.currentPlayer
}

func (g *Game) GetPlayerPiece(player game.Player) game.Piece {
	return Piece{
		player: player,
	}
}

func (g *Game) ExecuteAction(action game.IAction) {
	validAction, _ := g.rules.IsValidAction(action, g.currentPlayer, g.gameBoard)
	if !validAction {
		panic("Invalid action.")
	}

	switch typedAction := action.(type) {
	case *game.PlacePieceAction:
		g.gameBoard.PlacePiece(typedAction.X, typedAction.Y, typedAction.Piece)
	default:
		panic("Invalid action.")
	}

	gameOverState, _ := g.rules.IsGameOver(g.gameBoard)
	if gameOverState == game.NotGameOver {
		g.currentPlayer = (g.currentPlayer + 1) % 2
	}
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
