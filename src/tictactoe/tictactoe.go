package tictactoe

import (
	"github.com/chippolot/game-forge/src/game"
)

// Rules concrete implementation of the tic-tac-toe game rules
type Rules struct{}

func (r *Rules) IsValidMove(board game.IBoard, x, y int, piece game.Piece) bool {
	validSpace := x >= 0 && y >= 0 && x < board.GetWidth() && y < board.GetHeight()
	if !validSpace {
		return false
	}
	return board.GetPiece(x, y) == game.Empty
}

func (r *Rules) IsGameOver(board game.IBoard, piece game.Piece) bool {
	// Implement logic to check for a win or a draw
	// For tic-tac-toe, you would check for rows, columns, and diagonals
	return false
}

func (r *Rules) GetWinner(board game.IBoard) game.Piece {
	// Implement logic to determine the winner
	// For tic-tac-toe, you would check for rows, columns, and diagonals
	return game.Empty
}

// Game concrete implementation of the tic-tac-toe game
type Game struct {
	gameBoard    game.IBoard
	rules        game.IRules
	currentPiece game.Piece
}

func NewGame(board game.IBoard, rules game.IRules) game.IGame {
	return &Game{
		gameBoard:    board,
		rules:        rules,
		currentPiece: game.X, // Start with player X
	}
}

func (g *Game) GetName() string {
	return "Tic-Tac-Toe"
}

func (g *Game) GetDescription() string {
	return "Tic-tac-toe is a classic two-player game played on a 3x3 grid. Players take turns marking spaces with their respective symbols, typically \"X\" and \"O\", with the objective of placing three of their symbols in a row, column, or diagonal. The first player to achieve this goal wins the game. If all spaces are filled without a winner, the game ends in a draw. Tic-tac-toe is easy to learn, yet offers strategic depth, making it a timeless and engaging pastime for players of all ages."
}

func (g *Game) Start() {
	// Initialize the game
}

func (g *Game) GetCurrentPiece() game.Piece {
	return g.currentPiece
}

func (g *Game) MakeMove(x, y int) {
	if !g.rules.IsValidMove(g.gameBoard, x, y, g.currentPiece) {
		panic("Invalid move.")
	}
	g.gameBoard.PlacePiece(x, y, g.currentPiece)

	if g.rules.IsGameOver(g.gameBoard, g.currentPiece) {
		// Game over logic
	} else {
		// Switch player
		if g.currentPiece == game.X {
			g.currentPiece = game.O
		} else {
			g.currentPiece = game.X
		}
	}
}

func (g *Game) Restart() {
	g.gameBoard.Clear()
	g.currentPiece = game.X // Reset to player X
}
