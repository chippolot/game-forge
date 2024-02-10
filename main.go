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

	switch choice {
	case 1:
		gameInstance = tictactoe.NewGame()
	case 2:
		//gameInstance = gridlock.NewGame()
		return
	}

	gameInstance.Start()

	actionParser := game.NewActionParser()
	gameInstance.RegisterActions(actionParser)

	for {
		utils.ClearScreen()

		fmt.Println(gameInstance.GetName())
		fmt.Println("---------------------------------------")
		fmt.Println(utils.WrapLines(gameInstance.GetDescription(), 80))
		fmt.Println("---------------------------------------")
		fmt.Println()

		board := gameInstance.GetBoard()
		board.Print()

		fmt.Printf("Player %v's turn\n", (gameInstance.GetCurrentPlayer() + 1))
		fmt.Println("Enter action:")

		input, _ := reader.ReadString('\n')
		action, err := actionParser.ParseAction(input, gameInstance)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Scanln()
			continue
		}

		player := gameInstance.GetCurrentPlayer()
		if !gameInstance.GetRules().IsValidAction(action, player, board) {
			fmt.Println("Invalid move. Try again.")
			fmt.Scanln()
			continue
		}

		gameInstance.ExecuteAction(action)

		gameOverState, winningPlayer := gameInstance.GetRules().IsGameOver(board)
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
