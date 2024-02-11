package game

type ILogic interface {
	RegisterActions(actionParser *ActionParser)
	ExecuteAction(action IAction, state IGameState) (GameResult, error)
}
