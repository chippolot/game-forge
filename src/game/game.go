package game

type IGame interface {
	GetName() string
	GetDescription() string

	GetBoard() IBoard
	GetRules() IRules

	GetCurrentPlayer() Player
	GetPlayerPiece(player Player) Piece

	Start()
	RegisterActions(actionParser *ActionParser)
	ExecuteAction(action IAction)
	Restart()
}

type GameOverState int

const (
	NotGameOver GameOverState = iota
	GameWon
	GameTie
)
