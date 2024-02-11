package game

import (
	"fmt"

	"github.com/chippolot/game-forge/utils"
)

type IGameRenderer interface {
	Print(gameInstance IGame)
}

type SimpleGameRenderer struct {
}

func (r *SimpleGameRenderer) Print(gameInstance IGame) {
	// Print game metadata
	fmt.Println(gameInstance.GetName())
	fmt.Println("---------------------------------------")
	fmt.Println(utils.WrapLines(gameInstance.GetDescription(), 80))
	fmt.Println("---------------------------------------")
	fmt.Println()

	// Print board
	board := gameInstance.GetState().GetBoard()
	for y := 0; y < board.GetHeight(); y++ {
		for x := 0; x < board.GetWidth(); x++ {
			piece := board.GetPiece(x, y)
			if piece == nil {
				fmt.Print("- ")
			} else {
				fmt.Print(piece.GetDisplayString() + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// Print action prompt
	fmt.Printf("Player %v's turn\n", (gameInstance.GetState().GetCurrentPlayer() + 1))
	fmt.Print("Enter action: ")
}
