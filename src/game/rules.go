package game

type IRules interface {
	IsValidAction(action IAction, player Player, board IBoard) bool
	IsGameOver(board IBoard) (GameOverState, Player)
}
