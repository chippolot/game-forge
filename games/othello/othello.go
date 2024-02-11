package othello

import (
	"fmt"

	"github.com/chippolot/game-forge/game"
)

type Logic struct{}

func (l *Logic) RegisterActions(actionParser *game.ActionParser) {
	placePieceActionRegistration := game.ActionRegistration{
		Desc: game.ActionDesc{
			Keyword:     "place",
			Usage:       "place <coord>",
			Description: "Place a piece on the board at coordinates <coord>, flipping opponent's pieces according to Othello rules.",
		},
		ParseFunc: func(args []string, gameInstance game.IGame) (game.IAction, error) {
			return game.ParsePlacePieceAction(args, func() game.Piece {
				return getPlayerPiece(gameInstance.GetState().GetCurrentPlayer())
			})
		},
	}
	actionParser.RegisterAction(placePieceActionRegistration)
}

func (l *Logic) ExecuteAction(action game.IAction, state game.IGameState) (game.GameResult, error) {
	validAction, err := isValidAction(action, state)
	if !validAction {
		return game.GameResult{}, err
	}

	switch typedAction := action.(type) {
	case *game.PlacePieceAction:
		flipPieces(typedAction.X, typedAction.Y, typedAction.Piece, state.GetBoard())
	default:
		panic("Invalid action.")
	}

	updateScores(state)

	gameOverState, winningPlayer := isGameOver(state.GetBoard())
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
	case *game.PlacePieceAction:
		if canPlacePiece(typedAction.X, typedAction.Y, typedAction.Piece, state.GetBoard()) {
			return true, nil
		}
		return false, fmt.Errorf("invalid move")
	default:
		return false, fmt.Errorf("unsupported action %T", typedAction)
	}
}

func isGameOver(board game.IBoard) (game.GameResultState, game.Player) {
	if isBoardFullOrNoMoves(board) {
		return game.GameWon, getWinningPlayer(board)
	}
	return game.NotGameOver, 0
}

func isBoardFullOrNoMoves(board game.IBoard) bool {
	full := true
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < board.GetHeight(); y++ {
			if board.GetPiece(x, y) == nil {
				full = false
				if canPlacePiece(x, y, getPlayerPiece(0), board) || canPlacePiece(x, y, getPlayerPiece(1), board) {
					return false
				}
			}
		}
	}
	return full
}

func getWinningPlayer(board game.IBoard) game.Player {
	counts := map[game.Player]int{0: 0, 1: 0}
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < board.GetHeight(); y++ {
			if piece := board.GetPiece(x, y); piece != nil {
				counts[piece.GetPlayer()]++
			}
		}
	}
	if counts[0] > counts[1] {
		return 0
	}
	return 1
}

func canPlacePiece(x, y int, piece game.Piece, board game.IBoard) bool {
	if board.GetPiece(x, y) != nil {
		return false
	}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if checkDirectionForFlips(x, y, dx, dy, piece, board) {
				return true
			}
		}
	}
	return false
}

func flipPieces(x, y int, piece game.Piece, board game.IBoard) {
	board.PlacePiece(x, y, piece)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if checkDirectionForFlips(x, y, dx, dy, piece, board) {
				flipDirection(x, y, dx, dy, piece, board)
			}
		}
	}
}

func checkDirectionForFlips(x, y, dx, dy int, piece game.Piece, board game.IBoard) bool {
	x += dx
	y += dy
	opponentFound := false
	for board.IsInBounds(x, y) {
		currentPiece := board.GetPiece(x, y)
		if currentPiece == nil {
			return false
		}
		if currentPiece.GetPlayer() == piece.GetPlayer() {
			return opponentFound
		}
		opponentFound = true
		x += dx
		y += dy
	}
	return false
}

func flipDirection(x, y, dx, dy int, piece game.Piece, board game.IBoard) {
	x += dx
	y += dy
	for board.IsInBounds(x, y) && board.GetPiece(x, y).GetPlayer() != piece.GetPlayer() {
		board.PlacePiece(x, y, piece)
		x += dx
		y += dy
	}
}

type GameState struct {
	*game.CommonGameState
}

func (s *GameState) Reset() {
	s.CommonGameState.Reset()
	b := s.GetBoard()
	b.PlacePiece(3, 3, Piece{0})
	b.PlacePiece(4, 4, Piece{0})
	b.PlacePiece(3, 4, Piece{1})
	b.PlacePiece(4, 3, Piece{1})
	updateScores(s)
}

func updateScores(state game.IGameState) {
	scores := [2]int{0, 0}
	board := state.GetBoard()
	for x := 0; x < board.GetWidth(); x++ {
		for y := 0; y < board.GetHeight(); y++ {
			piece := board.GetPiece(x, y)
			if piece != nil {
				scores[piece.GetPlayer()] += 1
			}
		}
	}
	for i := 0; i < 2; i++ {
		state.SetPlayerScore(game.Player(i), scores[i])
	}
}

func NewState(board game.IBoard) game.IGameState {
	return &GameState{game.NewCommonGameState(board)}
}

type Piece struct {
	player game.Player
}

func NewGame(parser *game.ActionParser) game.IGame {
	name := "Othello"
	desc := "Othello, also known as Reversi, is a strategy board game for two players, played on an 8x8 uncheckered board. " +
		"There are sixty-four identical game pieces called disks, which are light on one side and dark on the other. " +
		"Players take turns placing disks on the board with their assigned color facing up. During a play, any disks of the opponent's color that are in a straight line and bounded by the disk just placed and another disk of the current player's color are turned over to the current player's color. " +
		"The objective of the game is to have the majority of disks turned to display your color when the last playable empty square is filled."
	metadata := game.NewMetadata(name, desc)
	logic := &Logic{}
	board := game.NewBoard(8, 8)
	state := NewState(board)
	renderer := &game.SimpleGameRenderer{PrintScores: true}
	logic.RegisterActions(parser)
	return game.NewGame(logic, state, renderer, metadata, parser)
}

func (p Piece) GetPlayer() game.Player {
	return p.player
}

func (p Piece) GetDisplayString() string {
	if p.player == 0 {
		return "B" // Black
	} else if p.player == 1 {
		return "W" // White
	}
	panic("unknown piece")
}

func getPlayerPiece(player game.Player) game.Piece {
	return Piece{player: player}
}
