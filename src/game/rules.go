package game

type IRules interface {
	IsValidAction(action IAction, player Player, board IBoard) (bool, error)
	IsGameOver(board IBoard) (GameOverState, Player)
}
