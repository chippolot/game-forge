package game

import (
	"fmt"

	"github.com/chippolot/game-forge/color"
	"github.com/chippolot/game-forge/utils"
)

type IGameRenderer interface {
	Print(gameInstance IGame)
}

type SimpleGameRenderer struct {
	PrintScores bool
}

func (r *SimpleGameRenderer) Print(gameInstance IGame) {
	// Print game metadata
	fmt.Println(gameInstance.GetName())
	fmt.Println("---------------------------------------")
	fmt.Println(utils.WrapLines(gameInstance.GetDescription(), 80))
	fmt.Println("---------------------------------------")
	fmt.Println()

	// (optionally) Print scores
	if r.PrintScores {
		printScore(0, gameInstance.GetState())
		printScore(1, gameInstance.GetState())
		fmt.Println()
	}

	// Print board
	board := gameInstance.GetState().GetBoard()
	printBoard(board)

	// Print action prompt
	fmt.Printf("Player %v's turn\n", (gameInstance.GetState().GetCurrentPlayer() + 1))
	fmt.Print("Enter action: ")
}

func printScore(player Player, state IGameState) {
	fmt.Printf("Player %v: %v\n", player, state.GetPlayerScore(player))
}

func printBoard(board IBoard) {
	// Print column coords
	fmt.Print("  ")
	for x := 0; x < board.GetWidth(); x++ {
		fmt.Printf("%c ", 'a'+x)
	}
	fmt.Println()

	// Print board squares and pieces
	height := board.GetHeight()
	for y := 0; y < height; y++ {
		// Print row coord
		fmt.Printf("%v ", y+1)
		for x := 0; x < board.GetWidth(); x++ {
			piece := board.GetPiece(x, y)
			if piece == nil {
				fmt.Print("- ")
			} else {
				c := color.Green
				if piece.GetPlayer() == 1 {
					c = color.Red
				}
				fmt.Print(c + piece.GetDisplayString() + color.Reset + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
