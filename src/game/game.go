package game

type IGame interface {
	GetName() string
	GetDescription() string

	Start()
	MakeMove(x, y int)
	GetCurrentPiece() Piece
	Restart()
}
