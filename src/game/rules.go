package game

type IRules interface {
	IsValidMove(board IBoard, x, y int, player Player, piece Piece) bool
	IsGameOver(board IBoard) (GameOverState, Player)
}
