package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chippolot/game-forge/src/game"
	tictactoe "github.com/chippolot/game-forge/src/games/tic-tac-toe"
	"github.com/chippolot/game-forge/src/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Game Hub!")
	fmt.Println("Please select a game to play:")
	fmt.Println("1. Tic-Tac-Toe")
	fmt.Println("2. Grid Lock")

	fmt.Print("Enter your choice: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid choice. Please enter 1 or 2.")
		return
	}

	var gameInstance game.IGame
	actionParser := game.NewActionParser()

	switch choice {
	case 1:
		gameInstance = tictactoe.NewGame(actionParser)
	case 2:
		//gameInstance = gridlock.NewGame()
		return
	}

	gameInstance.Start()

	for {
		utils.ClearScreen()

		// Print game metadata
		fmt.Println(gameInstance.GetMetadata().Name)
		fmt.Println("---------------------------------------")
		fmt.Println(utils.WrapLines(gameInstance.GetMetadata().Decription, 80))
		fmt.Println("---------------------------------------")
		fmt.Println()

		// Print board
		gameInstance.Print()

		// Print action prompt
		fmt.Printf("Player %v's turn\n", (gameInstance.GetCurrentPlayer() + 1))
		fmt.Println("Enter action:")

		// Parse player action
		input, _ := reader.ReadString('\n')
		action, err := actionParser.ParseAction(input, gameInstance)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Scanln()
			continue
		}

		// Execute player action
		gameState, err := gameInstance.ExecuteAction(action)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Scanln()
			continue
		}

		// Check for game over
		if gameState.State == game.GameWon {
			fmt.Printf("Player %v Wins!\n", gameState.WinningPlayer)
			break
		}
		if gameState.State == game.GameTie {
			fmt.Println("Tied game.")
			break
		}
	}
}
