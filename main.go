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
	game := tictactoe.NewGame(board, rules)

	game.Start()

	for {
		utils.ClearScreen()

		fmt.Println(game.GetName())
		fmt.Println("---------------------------------------")
		fmt.Println(utils.WrapLines(game.GetDescription(), 80))
		fmt.Println("---------------------------------------")
		fmt.Println()

		board.Print()

		fmt.Printf("Player %v's turn (X/O)\n", game.GetCurrentPiece())
		fmt.Println("Enter row and column separated by space (e.g., '1 1' for top-left corner):")

		var row, col int
		fmt.Scanln(&row, &col)

		x := col - 1
		y := row - 1

		if !rules.IsValidMove(board, x, y, game.GetCurrentPiece()) {
			fmt.Println("Invalid move. Try again.")
			fmt.Scanln()
			continue
		}

		game.MakeMove(x, y)
	}
}
