package game

type IGame interface {
	IMetadata

	GetCurrentPlayer() Player

	Start()
	ExecuteAction(action IAction) (GameResult, error)
	Print()
	Restart()
}

type GameResult struct {
	State         GameResultState
	WinningPlayer Player
}

type GameResultState int

const (
	NotGameOver GameResultState = iota
	GameWon
	GameTie
)
