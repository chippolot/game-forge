package checkers

import (
	"fmt"

	"github.com/chippolot/game-forge/game"
)

type Logic struct{}

func (l *Logic) RegisterActions(actionParser *game.ActionParser) {
	movePieceActionRegistration := game.ActionRegistration{
		Desc: game.MovePieceActionDesc(),
		ParseFunc: func(args []string, gameInstance game.IGame) (game.IAction, error) {
			return game.ParseMovePieceAction(args, -1)
		},
	}
	actionParser.RegisterAction(movePieceActionRegistration)
}

func (l *Logic) ExecuteAction(action game.IAction, state game.IGameState) (game.GameResult, error) {
	validAction, err := isValidAction(action, state)
	if !validAction {
		return game.GameResult{}, err
	}

	switch typedAction := action.(type) {
	case *game.MovePieceAction:
		executeMove(typedAction.Start, typedAction.Moves[0], state)
	default:
		panic("Invalid action.")
	}

	gameOverState, winningPlayer := isGameOver(state)
	if gameOverState == game.NotGameOver {
		currentPlayer := state.GetCurrentPlayer()
		currentPlayer = (currentPlayer + 1) % 2
		state.SetCurrentPlayer(currentPlayer)
	}

	gameResult := game.GameResult{
		State:         gameOverState,
		WinningPlayer: winningPlayer,
	}
	return gameResult, nil
}

func isValidAction(action game.IAction, state game.IGameState) (bool, error) {
	switch typedAction := action.(type) {
	case *game.MovePieceAction:
		return isValidMove(typedAction.Start, typedAction.Moves[0], state), fmt.Errorf("invalid move")
	default:
		return false, fmt.Errorf("unsupported action %T", typedAction)
	}
}

func isGameOver(state game.IGameState) (game.GameResultState, game.Player) {
	if noMoreMoves(state) {
		return game.GameWon, calculateWinningPlayer(state)
	}
	return game.NotGameOver, 0
}

func noMoreMoves(state game.IGameState) bool {
	// Simplified check; in a real game, you'd check for all possible moves for the current player
	board := state.GetBoard()
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < board.GetHeight(); y++ {
			piece := board.GetPiece(x, y)
			if piece != nil && piece.GetPlayer() == state.GetCurrentPlayer() {
				// Check for possible moves for this piece
				// This is a simplified check; a real check would consider multiple jump moves and king movement
				for dx := -1; dx <= 1; dx += 2 {
					for dy := -1; dy <= 1; dy += 2 {
						if isValidMove(game.Coord{X: x, Y: y}, game.Coord{X: x + dx, Y: y + dy}, state) {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func calculateWinningPlayer(state game.IGameState) game.Player {
	// Count pieces for each player
	board := state.GetBoard()
	counts := make(map[game.Player]int)
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < board.GetHeight(); y++ {
			piece := board.GetPiece(x, y)
			if piece != nil {
				counts[piece.GetPlayer()]++
			}
		}
	}

	// Determine the player with more pieces
	if counts[0] > counts[1] {
		return 0
	} else if counts[1] > counts[0] {
		return 1
	}
	// If equal or no pieces, this logic would need to decide based on the rules (e.g., stalemate rules)
	return -1 // Indicate a draw or unresolved game state
}

func isValidMove(from game.Coord, to game.Coord, state game.IGameState) bool {
	board := state.GetBoard()
	fromPiece := board.GetPiece(from.X, from.Y)
	if fromPiece == nil || fromPiece.GetPlayer() != state.GetCurrentPlayer() {
		return false
	}

	// Check if the target location is within bounds and unoccupied
	if !board.IsInBounds(to.X, to.Y) || board.GetPiece(to.X, to.Y) != nil {
		return false
	}

	// Check for a valid move distance (simple move or capture move)
	dx := to.X - from.X
	dy := to.Y - from.Y
	if fromPiece.(*Piece).IsKing() {
		// Kings can move backward
		if (abs(dx) != 1 || abs(dy) != 1) && (abs(dx) != 2 || abs(dy) != 2) {
			return false
		}
	} else {
		// Regular pieces can only move forward
		if state.GetCurrentPlayer() == 0 && dy <= 0 {
			return false
		}
		if state.GetCurrentPlayer() == 1 && dy >= 0 {
			return false
		}
		if (abs(dx) != 1 || abs(dy) != 1) && (abs(dx) != 2 || abs(dy) != 2) {
			return false
		}
	}

	// Check for capture
	if abs(dx) == 2 && abs(dy) == 2 {
		midX := from.X + dx/2
		midY := from.Y + dy/2
		midPiece := board.GetPiece(midX, midY)
		if midPiece == nil || midPiece.GetPlayer() == state.GetCurrentPlayer() {
			return false
		}
	}

	return true
}

func executeMove(from game.Coord, to game.Coord, state game.IGameState) {
	board := state.GetBoard()
	dx := to.X - from.X
	dy := to.Y - from.Y
	board.MovePiece(from.X, from.Y, to.X, to.Y)
	if abs(dx) == 2 && abs(dy) == 2 {
		// Capture move
		midX := from.X + dx/2
		midY := from.Y + dy/2
		board.RemovePiece(midX, midY)
	}

	// Check for promotion to King
	if (state.GetCurrentPlayer() == 0 && to.Y == board.GetHeight()-1) || (state.GetCurrentPlayer() == 1 && to.Y == 0) {
		piece := board.GetPiece(to.X, to.Y).(*Piece)
		piece.king = true
	}
}

type GameState struct {
	*game.CommonGameState
}

func NewState(board game.IBoard) game.IGameState {
	return &GameState{game.NewCommonGameState(board)}
}

func (s *GameState) Reset() {
	s.CommonGameState.Reset()
	initializeBoard(s.GetBoard())
}

func initializeBoard(board game.IBoard) {
	// Place pieces on the board in their initial positions
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < 3; y++ {
			if (x+y)%2 == 1 {
				board.PlacePiece(x, y, &Piece{player: 1})
			}
		}
		for y := board.GetHeight() - 3; y < board.GetHeight(); y++ {
			if (x+y)%2 == 1 {
				board.PlacePiece(x, y, &Piece{player: 0})
			}
		}
	}
}

type Piece struct {
	player game.Player
	king   bool
}

func (p Piece) GetPlayer() game.Player {
	return p.player
}

func (p Piece) IsKing() bool {
	return p.king
}

func (p Piece) GetDisplayString() string {
	if p.player == 0 {
		return "B" // Black
	} else if p.player == 1 {
		return "W" // White
	}
	panic("unknown piece")
}

func NewGame(parser *game.ActionParser) game.IGame {
	name := "Checkers"
	desc := "Checkers is a strategy board game for two players which involves diagonal moves of uniform game pieces and mandatory captures by jumping over opponent pieces."
	metadata := game.NewMetadata(name, desc)
	logic := &Logic{}
	board := game.NewBoard(8, 8)
	state := NewState(board)
	renderer := &game.SimpleGameRenderer{PrintScores: true}
	logic.RegisterActions(parser)
	return game.NewGame(logic, state, renderer, metadata, parser)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
