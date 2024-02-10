package game

type IGame interface {
	GetName() string
	GetDescription() string

	Start()
	MakeMove(x, y int, piece Piece)
	GetCurrentPlayer() Player
	GetPlayerPiece(player Player) Piece
	Restart()
}

type GameOverState int

const (
	NotGameOver GameOverState = iota
	GameWon
	GameTie
)
