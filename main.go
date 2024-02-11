package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chippolot/game-forge/game"
	"github.com/chippolot/game-forge/games/othello"
	tictactoe "github.com/chippolot/game-forge/games/tic-tac-toe"
	"github.com/chippolot/game-forge/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Game Hub!")
	fmt.Println("Please select a game to play:")
	fmt.Println("1. Tic-Tac-Toe")
	fmt.Println("2. Othello")

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
		gameInstance = othello.NewGame(actionParser)
	}

	gameInstance.Start()

	for {
		utils.ClearScreen()

		// Print game instance
		gameInstance.Print()

		// Parse player action
		input, _ := reader.ReadString('\n')
		action, err := actionParser.ParseAction(input, gameInstance)
		if err != nil {
			fmt.Println("Error:", err)
			actionParser.PrintAvailableActions()
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
