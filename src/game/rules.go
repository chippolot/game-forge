package game

type IRules interface {
	IsValidMove(board IBoard, x, y int, piece Piece) bool
	IsGameOver(board IBoard, piece Piece) bool
	GetWinner(board IBoard) Piece
}
