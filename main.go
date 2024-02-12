package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chippolot/game-forge/game"
	"github.com/chippolot/game-forge/games/checkers"
	"github.com/chippolot/game-forge/games/othello"
	"github.com/chippolot/game-forge/games/tictactoe"
	"github.com/chippolot/game-forge/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Game Hub!")
	fmt.Println("Please select a game to play:")
	fmt.Println("1. Tic-Tac-Toe")
	fmt.Println("2. Othello")
	fmt.Println("3. Checkers")

	fmt.Print("Enter your choice: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || (choice < 1 || choice > 3) {
		fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
		return
	}

	var gameInstance game.IGame
	actionParser := game.NewActionParser()

	switch choice {
	case 1:
		gameInstance = tictactoe.NewGame(actionParser)
	case 2:
		gameInstance = othello.NewGame(actionParser)
	case 3:
		gameInstance = checkers.NewGame(actionParser)
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
