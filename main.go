package main

import (
	"fmt"

	"github.com/chippolot/game-forge/src/game"
	"github.com/chippolot/game-forge/src/tictactoe"
	"github.com/chippolot/game-forge/src/utils"
)

func main() {
	board := game.NewBoard(3, 3)
	rules := &tictactoe.Rules{}
	gameInstace := tictactoe.NewGame(board, rules)

	gameInstace.Start()

	for {
		utils.ClearScreen()

		fmt.Println(gameInstace.GetName())
		fmt.Println("---------------------------------------")
		fmt.Println(utils.WrapLines(gameInstace.GetDescription(), 80))
		fmt.Println("---------------------------------------")
		fmt.Println()

		board.Print()

		fmt.Printf("Player %v's turn\n", (gameInstace.GetCurrentPlayer() + 1))
		fmt.Println("Enter row and column separated by space (e.g., '1 1' for top-left corner):")

		var row, col int
		fmt.Scanln(&row, &col)

		x := col - 1
		y := row - 1

		player := gameInstace.GetCurrentPlayer()
		piece := gameInstace.GetPlayerPiece(player)
		if !rules.IsValidMove(board, x, y, player, piece) {
			fmt.Println("Invalid move. Try again.")
			fmt.Scanln()
			continue
		}

		gameInstace.MakeMove(x, y, piece)

		gameOverState, winningPlayer := rules.IsGameOver(board)
		if gameOverState == game.GameWon {
			fmt.Printf("Player %v Wins!\n", winningPlayer)
			break
		}
		if gameOverState == game.GameTie {
			fmt.Println("Tied game.")
			break
		}
	}
}
