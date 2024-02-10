package game

type IGame interface {
	GetName() string
	GetDescription() string

	GetBoard() IBoard

	GetCurrentPlayer() Player
	GetPlayerPiece(player Player) Piece

	Start()
	RegisterActions(actionParser *ActionParser)
	ExecuteAction(action IAction) (GameState, error)
	Restart()
}

type GameState struct {
	State         GameOverState
	WinningPlayer Player
}

type GameOverState int

const (
	NotGameOver GameOverState = iota
	GameWon
	GameTie
)
