package tictactoe

import (
	"github.com/chippolot/game-forge/src/game"
)

// Rules concrete implementation of the tic-tac-toe game rules
type Rules struct{}

func (r *Rules) IsValidMove(board game.IBoard, x, y int, player game.Player, piece game.Piece) bool {
	validSpace := x >= 0 && y >= 0 && x < board.GetWidth() && y < board.GetHeight()
	if !validSpace {
		return false
	}
	return board.GetPiece(x, y) == nil
}

func (r *Rules) IsGameOver(board game.IBoard) (game.GameOverState, game.Player) {
	hasWinner, winningPlayer := getWinner(board)
	if hasWinner {
		return game.GameWon, winningPlayer
	}
	if isBoardFilled(board) {
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

func isBoardFilled(board game.IBoard) bool {
	for col := 0; col < board.GetWidth(); col++ {
		for row := 0; row < board.GetHeight(); row++ {
			if board.GetPiece(col, row) == nil {
				return false
			}
		}
	}
	return true
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

func NewGame(board game.IBoard, rules game.IRules) game.IGame {
	return &Game{
		gameBoard:     board,
		rules:         rules,
		currentPlayer: 0,
	}
}

func (g *Game) GetName() string {
	return "Tic-Tac-Toe"
}

func (g *Game) GetDescription() string {
	return "Tic-tac-toe is a classic two-player game played on a 3x3 grid. Players take turns marking spaces with their respective symbols, typically \"X\" and \"O\", with the objective of placing three of their symbols in a row, column, or diagonal. The first player to achieve this goal wins the game. If all spaces are filled without a winner, the game ends in a draw. Tic-tac-toe is easy to learn, yet offers strategic depth, making it a timeless and engaging pastime for players of all ages."
}

func (g *Game) Start() {
	g.currentPlayer = 0
}

func (g *Game) GetCurrentPlayer() game.Player {
	return g.currentPlayer
}

func (g *Game) GetPlayerPiece(player game.Player) game.Piece {
	if player == 0 {
		return XPiece{
			player: player,
		}
	} else if player == 1 {
		return OPiece{
			player: player,
		}
	}
	return nil
}

func (g *Game) MakeMove(x, y int, piece game.Piece) {
	if !g.rules.IsValidMove(g.gameBoard, x, y, g.currentPlayer, piece) {
		panic("Invalid move.")
	}
	g.gameBoard.PlacePiece(x, y, piece)

	gameOverState, _ := g.rules.IsGameOver(g.gameBoard)
	if gameOverState == game.NotGameOver {
		g.currentPlayer = (g.currentPlayer + 1) % 2
	}
}

func (g *Game) Restart() {
	g.gameBoard.Clear()
	g.Start()
}

type XPiece struct {
	player game.Player
}

func (p XPiece) GetPlayer() game.Player {
	return p.player
}

func (p XPiece) GetDisplayString() string {
	return "x"
}

type OPiece struct {
	player game.Player
}

func (p OPiece) GetPlayer() game.Player {
	return p.player
}

func (p OPiece) GetDisplayString() string {
	return "o"
}
